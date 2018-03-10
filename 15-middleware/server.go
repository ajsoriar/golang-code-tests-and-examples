// Instructions: go run server.go

package main

import (
	"fmt"
	"net/http"
	"path"
)

//var AccessControlAllowOrigin = `"*"`
//var AccessControlAllowOrigin = "http://localhost:8080"

func main() {
	http.HandleFunc("/login", _response_login)
	http.HandleFunc("/user", _response_user)
	http.HandleFunc("/", _welcome)
	http.HandleFunc("/shutdown", _shutdown)
	http.ListenAndServe(":7009", nil)
}

// -----------------------------------------------
// - WELCOME
// -----------------------------------------------

func _welcome(w http.ResponseWriter, r *http.Request) {
	js := "<b>Hello HTML (1)! This is a go server!</b>"
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "A Go Web Server")
	//w.Header().Set("Access-Control-Allow-Origin", AccessControlAllowOrigin)
	w.Write([]byte(js))
}

// -----------------------------------------------
// - SHUTDOWN
// -----------------------------------------------

// Check out this to stop the server: https://stackoverflow.com/questions/39320025/how-to-stop-http-listenandserve

func _shutdown(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Shutdown!")
	//http.Shutdown(nil)
}

// -----------------------------------------------
// - MORE
// -----------------------------------------------

func _response_login(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("data/atresplayer/login", "response.json")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "A Go Web Server")
	//w.Header().Set("Access-Control-Allow-Origin", AccessControlAllowOrigin)
	http.ServeFile(w, r, fp)
}

func _response_user(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("data/atresplayer/login", "response.json")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "A Go Web Server")
	//w.Header().Set("Access-Control-Allow-Origin", AccessControlAllowOrigin)
	http.ServeFile(w, r, fp)
}
