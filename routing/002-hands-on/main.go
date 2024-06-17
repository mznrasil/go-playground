package main

import (
	"io"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "<h1>Hello from index route</h1>")
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "<h1>Hello from dog route</h1>")
}

func me(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "<h1>Hello from Rasil</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
