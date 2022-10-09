package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	initialProducts := []Product{
		{Name: "T-shirt", Description: "Coolest t-shirt", Price: 19.0, Quantity: 5},
		{"Headset", "Comfortable", 10.0, 5},
		{Name: "Hat", Description: "Coolest hat", Price: 9.5, Quantity: 50},
	}

	temp.ExecuteTemplate(w, "Index", initialProducts)
}
