/*

- Run the example this way:

    http://localhost:8000/givemedata

- In the browser you will see a json

- The header of the response will be: 

    Content-Type:application/json
    Andres-Header:Andres header!!!

*/

package main

import (
    "io"
    "net/http"
    "fmt"
    "os"
    //"encoding/json"
    "io/ioutil"
)

var counter int

func main() {

    fmt.Printf("%s\n", "main()" )
	
	http.HandleFunc("/", hello)
	http.HandleFunc("/andres", helloAndres )
	http.HandleFunc("/givemedata", givemedata )

	http.ListenAndServe(":8000", nil)

    fmt.Printf("%s\n", "Server started in http://localhost:8000/" )

    counter++
}

func hello(w http.ResponseWriter, r *http.Request) {

    fmt.Printf("%s\n", "hello()" )

	io.WriteString(w, "Hello world!")

    counter++
}

func helloAndres(w http.ResponseWriter, r *http.Request) {

    fmt.Printf("%s\n", "helloAndres()" )

	io.WriteString(w, "Hello andres!")

    counter++
}

func givemedata(w http.ResponseWriter, r *http.Request) {

    fmt.Printf("%s\n", "givemedata()" )

    file, e := ioutil.ReadFile("./andres.json")
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }

    // Update the header in order to return "application/json"
    // Check out this: http://www.alexedwards.net/blog/golang-response-snippets

    w.Header().Set("Server", "A Go Web Server")
    w.Header().Set("Andres-Header", "Andres header!!!")
    //w.WriteHeader(200)
    w.Header().Set("Content-Type", "application/json")

    io.WriteString(w, string(file))

    counter++

    fmt.Printf( "%d\n",counter )
}