package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("views/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dogs/", dogs)
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

func index(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func dogs(w http.ResponseWriter, r *http.Request) {
	dogs := []string{
		"Labarador",
		"German Shepherd",
		"Pug",
		"Japanese",
	}

	data := struct {
		Dogs []string
	}{
		Dogs: dogs,
	}

	err := tmpl.ExecuteTemplate(w, "dogs.html", data)
	if err != nil {
		log.Fatalln(err)
	}
}

type User struct {
	Fname, Lname string
}

func me(w http.ResponseWriter, r *http.Request) {
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")

	user := User{
		Fname: fname,
		Lname: lname,
	}

	data := struct {
		User User
	}{
		User: user,
	}

	err := tmpl.ExecuteTemplate(w, "me.html", data)
	if err != nil {
		log.Fatalln(err)
	}
}
