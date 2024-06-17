package main

import (
	"net/http"
)

type MyHandler int

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	w.Header().Set("Name", "Rasil")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<h1>This is a random text</h1>"))
	// fmt.Fprintln(w, "<h1>This is a random text</h1>")
}

func main() {
	var myHandler MyHandler

	srv := &http.Server{
		Addr:    ":8080",
		Handler: myHandler,
	}

	srv.ListenAndServe()
}
