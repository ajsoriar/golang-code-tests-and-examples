package main

import (
	"html/template"
	"net/http"
)

func handler1(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/t1.html", "templates/t2.html")
	t.Execute(w, "Andres")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/t1.html", "templates/t2.html")
	t.ExecuteTemplate(w, "t2.html", "Golang")
}

func handler3(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/home.html")
	t.ExecuteTemplate(w, "home.html", "Hello world")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/index.html", handler3)
	http.HandleFunc("/t1", handler1)
	http.HandleFunc("/t2", handler2)

	server.ListenAndServe()
}
