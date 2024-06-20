package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.Handle("/pics/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.gohtml")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, nil)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title></title>
</head>
<body>

<h1>Pictures of dogs:</h1>
<img src="/pics/dog.jpeg">
<img src="/pics/dog1.jpeg">
<img src="/pics/dog2.jpeg">

</body>
</html>
