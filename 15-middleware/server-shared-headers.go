// Instructions: go run server.go

package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
)

//var AccessControlAllowOrigin = `"*"`
//var AccessControlAllowOrigin = "http://localhost:8080"

func main() {

	//lt1 := chainMiddleware(withLogging, withTracing)
	//lt2 := chainMiddleware(andres)
	lt3 := chainMiddleware(setAndresHeaders, withLogging, withTracing)

	http.HandleFunc("/login", lt3(_response_login))
	http.HandleFunc("/user", lt3(_response_user))
	http.HandleFunc("/shutdown", _shutdown)
	//http.HandleFunc("/", lt2(welcome))
	http.ListenAndServe(":7009", nil)
}

// Source here: https://hackernoon.com/simple-http-middleware-with-go-79a4ad62889b
// middleware provides a convenient mechanism for filtering HTTP requests
// entering the application. It returns a new handler which may perform various
// operations and should finish by calling the next HTTP handler.
type middleware func(next http.HandlerFunc) http.HandlerFunc

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Logged connection from %s", r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}

func withTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Tracing request for %s", r.RequestURI)
		next.ServeHTTP(w, r)
	}
}

func andres(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Andrés 1, 2 y 3!")
		next.ServeHTTP(w, r)
	}
}

func chainMiddleware(mw ...middleware) middleware {
	return func(final http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			last := final
			for i := len(mw) - 1; i >= 0; i-- {
				last = mw[i](last)
			}
			last(w, r)
		}
	}
}

// ----- COMMON HEADERS -------

func setAndresHeaders(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Andrés!!! setAndresHeaders()")

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Server", "A Go Web Server")
		w.Header().Set("Andres-Header", "Hello!!!")
		next.ServeHTTP(w, r)
	}
}

// -----------------------------------------------
// - WELCOME
// -----------------------------------------------

func welcome(w http.ResponseWriter, r *http.Request) {
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
	fp := path.Join("data", "response.json")
	/*
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Server", "A Go Web Server")
		//w.Header().Set("Access-Control-Allow-Origin", AccessControlAllowOrigin)
	*/
	http.ServeFile(w, r, fp)
}

func _response_user(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("data", "response.json")
	/*)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "A Go Web Server")
	//w.Header().Set("Access-Control-Allow-Origin", AccessControlAllowOrigin)
	*/

	http.ServeFile(w, r, fp)
}
