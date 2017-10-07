/* EXAMPLE 1 */

/*

package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}

*/

/* EXAMPLE 2 */

/*
package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	http.ListenAndServe(":8000", mux)
}
*/

/*

In the example above, we don’t use the `nil` in function `http.ListenAndServe` any more. Instead, 
replace it with a variable with type `*ServeMux`. As you may guess, this example does exactly the 
same thing as the previous example but we use the variable `mux` to register the handle function 
instead of directly registering from the `net/http` package. What’s going on underneath?  Well, 
the reason you can directly register the handle function in the package level is because `net/http` 
has a default `*ServeMux` inside the package, where now we defined our own.

*/

/* EXAMPLE 3 */

package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world 3!")
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{
		Addr:    ":8000",
		Handler: &myHandler{},
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = hello

	server.ListenAndServe()
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	io.WriteString(w, "My server: "+r.URL.String())
}

/*

To confirm your guess, this time we’re doing the same thing again, which is showing the `Hello world!` 
string on the screen. However, we not only define the `*ServeMux`, but also the variable server with 
type `http.Server`. At this point, you should know why we are able to run and serve the HTTP server 
directly from the `net/http` package. Yes, it has a default server inside the package as well. A new 
thing we see here is the type `myHandler` we defined and its method `ServeHTTP`. How is it possible 
to define a custom type and use it in an unmodified function from a standard library in a static 
programming language? The truth is simple. The `Handler` is an interface and only required to implement 
one method whose signature is `func(w http.ResponseWriter, r *http.Request)`(https://gowalker.org/net/http#Handler) 
and must be named as `ServeHTTP`. The type `myHandler` has the method that the interface expects, so we are good.

*/