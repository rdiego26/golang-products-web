package main

import (
	_ "github.com/lib/pq"
	"golang-products-web/models"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	var products = models.GetAllProducts()

	temp.ExecuteTemplate(w, "Index", products)
}
