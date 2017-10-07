/*

- Run the example this way: 

http://localhost:8000/givemedata

- In the browser you will see a json

- The header of the response will be: 

Content-Type:text/plain; charset=utf-8 
(we dont like this, we will overwrite the response in the next example)

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

func main() {
	
	http.HandleFunc("/", hello)
	http.HandleFunc("/andres", helloAndres )
	http.HandleFunc("/givemedata", givemedata )

	http.ListenAndServe(":8000", nil)
}


func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func helloAndres(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello andres!")
}

func givemedata(w http.ResponseWriter, r *http.Request) {

    file, e := ioutil.ReadFile("./andres.json")
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }

    io.WriteString(w, string(file))

    // fmt.Printf("%s\n", string(file))

    // //m := new(Dispatch)
    // //var m interface{}
    // var jsontype jsonobject
    // json.Unmarshal(file, &jsontype)
    // fmt.Printf("Results: %v\n", jsontype)
}