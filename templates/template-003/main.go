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

type item struct {
	Name  string
	Price float64
}

type meal struct {
	Meal  string
	Items []item
}

type menu []meal

type hotel struct {
	Name string
	Menu menu
}

type hotels []hotel

func main() {
	menu1 := menu{
		{
			Meal: "Breakfast",
			Items: []item{
				{"Oatmeal", 4.95},
				{"Cheerios", 3.95},
				{"Omelet", 6.95},
			},
		},
		{
			Meal: "Lunch",
			Items: []item{
				{"Burger", 5.95},
				{"Pizza", 7.95},
				{"Wrap", 6.95},
			},
		},
		{
			Meal: "Dinner",
			Items: []item{
				{"Steak", 12.95},
				{"Pasta", 8.95},
				{"Chicken", 9.95},
			},
		},
	}

	menu2 := menu{
		{
			Meal: "Breakfast",
			Items: []item{
				{"Biscuits and Gravy", 7.95},
				{"Pancakes", 5.95},
				{"French Toast", 6.95},
			},
		},
		{
			Meal: "Lunch",
			Items: []item{
				{"Soup", 3.95},
				{"Salad", 4.95},
				{"Sandwich", 4.95},
			},
		},
		{
			Meal: "Dinner",
			Items: []item{
				{"Shrimp", 14.95},
				{"Salmon", 13.95},
				{"Sushi", 11.95},
			},
		},
	}

	h := hotels{
		{
			Name: "Hotel California",
			Menu: menu1,
		},
		{
			Name: "Bates Motel",
			Menu: menu2,
		},
	}

	err := tmpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
