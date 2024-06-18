package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/photo.jpg", handlePhoto)
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
    <img src="/photo.jpg" />
  `)
}

func handlePhoto(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("photo.jpg")
	if err != nil {
		http.Error(w, "File Not Found", 404)
	}
	defer file.Close()

	io.Copy(w, file)
}
