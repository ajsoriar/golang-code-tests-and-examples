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

package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world 2!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	http.ListenAndServe(":8000", mux)
}

/*

In the example above, we don’t use the `nil` in function `http.ListenAndServe` any more. Instead, 
replace it with a variable with type `*ServeMux`. As you may guess, this example does exactly the 
same thing as the previous example but we use the variable `mux` to register the handle function 
instead of directly registering from the `net/http` package. What’s going on underneath?  Well, 
the reason you can directly register the handle function in the package level is because `net/http` 
has a default `*ServeMux` inside the package, where now we defined our own.

*/