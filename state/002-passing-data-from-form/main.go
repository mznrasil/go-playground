package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	subscribed := r.FormValue("subscribed")

	tmpl.Execute(w, struct {
		Fname      string
		Lname      string
		Subscribed string
	}{
		Fname:      fname,
		Lname:      lname,
		Subscribed: subscribed,
	})
}
