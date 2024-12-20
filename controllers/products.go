package controllers

import (
	"github.com/google/uuid"
	"golang-products-web/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, _ *http.Request) {
	var products = models.GetAllProducts()

	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, _ *http.Request) {

	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error during price conversion:", err)
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error during quantity conversion:", err)
		}

		models.CreateProduct(name, description, float64(convertedPrice), convertedQuantity)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	log.Printf("Received productId=%s\n", productId)

	models.DeleteProduct(productId)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	log.Printf("Received productId=%s\n", productId)

	parsedProductId, err := uuid.Parse(productId)
	if err != nil {
		panic(err.Error())
	}

	temp.ExecuteTemplate(w, "Edit", models.GetProduct(parsedProductId))
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		productId := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		log.Printf("Received from POST: %s %s %s %s %s\n", productId, name, description, price, quantity)

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error during price conversion:", err)
			panic(err.Error())
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error during quantity conversion:", err)
			panic(err.Error())
		}

		models.UpdateProduct(productId, name, description, convertedPrice, convertedQuantity)

		parsedProductId, err := uuid.Parse(productId)
		if err != nil {
			panic(err.Error())
		}

		temp.ExecuteTemplate(w, "Edit", models.GetProduct(parsedProductId))
	}
}
