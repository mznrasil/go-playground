package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.html"))
}

type CustomHandler int

func (h CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method  string
		Url     *url.URL
		Payload url.Values
	}{
		Method:  r.Method,
		Url:     r.URL,
		Payload: r.Form,
	}

	tmpl.Execute(w, data)
}

func main() {
	var handler CustomHandler
	http.ListenAndServe(":8080", handler)
}
