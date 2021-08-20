package main

import (
	"net/http"
	"log"
)

type hello struct{}

func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World</h1>"))
}

func main() {
	http.HandleFunc("/chapter1", func(w http.ResponseWriter, r *http.Request) {
		msg := "<h1>Chapter 1</h1>"
		w.Write([]byte(msg))
	})
	http.Handle("/", hello{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}