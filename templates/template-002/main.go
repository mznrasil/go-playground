package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.html"))
}

type Hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

type CaliforniaHotels struct {
	Hotels []Hotel
}

func main() {
	hotels := CaliforniaHotels{
		Hotels: []Hotel{
			{
				Name:    "Hotel California",
				Address: "1234 California St",
				City:    "San Francisco",
				Zip:     "94123",
				Region:  "Northern",
			},
			{
				Name:    "Hotel California",
				Address: "1234 California St",
				City:    "San Francisco",
				Zip:     "94123",
				Region:  "Northern",
			},
			{
				Name:    "Hotel California",
				Address: "1234 California St",
				City:    "San Francisco",
				Zip:     "94123",
				Region:  "Northern",
			},
		},
	}

	err := tmpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
}
