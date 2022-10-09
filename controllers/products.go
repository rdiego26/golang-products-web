package controllers

import (
	"golang-products-web/models"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, _ *http.Request) {
	var products = models.GetAllProducts()

	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, _ *http.Request) {

	temp.ExecuteTemplate(w, "New", nil)
}
