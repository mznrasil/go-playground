package main

import (
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	http.HandleFunc("/dog/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<h1>Hello Dog</h1>"))
	})

	http.HandleFunc("/cat", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<h1>Hello Cat</h1>"))
	})

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
