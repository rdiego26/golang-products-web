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

// Home page
func Index(w http.ResponseWriter, _ *http.Request) {
	products, err := models.GetAllProducts()
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		log.Printf("Error fetching products: %v\n", err)
		return
	}

	if err := temp.ExecuteTemplate(w, "Index", products); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error rendering template: %v\n", err)
	}
}

// New product page
func New(w http.ResponseWriter, _ *http.Request) {
	if err := temp.ExecuteTemplate(w, "New", nil); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error rendering template: %v\n", err)
	}
}

// Insert product
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			http.Error(w, "Invalid price", http.StatusBadRequest)
			log.Printf("Error converting price: %v\n", err)
			return
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			http.Error(w, "Invalid quantity", http.StatusBadRequest)
			log.Printf("Error converting quantity: %v\n", err)
			return
		}

		err = models.CreateProduct(name, description, convertedPrice, convertedQuantity)
		if err != nil {
			http.Error(w, "Error creating product", http.StatusInternalServerError)
			log.Printf("Error creating product: %v\n", err)
			return
		}

		log.Printf("Product successfully created: %s\n", name)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// Delete product
func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	if _, err := uuid.Parse(productId); err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Printf("Invalid ID: %v\n", err)
		return
	}

	err := models.DeleteProduct(productId)
	if err != nil {
		http.Error(w, "Error deleting product", http.StatusInternalServerError)
		log.Printf("Error deleting product: %v\n", err)
		return
	}

	log.Printf("Product successfully deleted: %s\n", productId)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Edit product
func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	if _, err := uuid.Parse(productId); err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Printf("Invalid ID: %v\n", err)
		return
	}

	product, err := models.GetProduct(productId)
	if err != nil {
		http.Error(w, "Error fetching product", http.StatusInternalServerError)
		log.Printf("Error fetching product: %v\n", err)
		return
	}

	if err := temp.ExecuteTemplate(w, "Edit", product); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error rendering template: %v\n", err)
	}
}

// Update product
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		productId := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		log.Printf("Received for update: %s %s %s %s %s\n", productId, name, description, price, quantity)

		if _, err := uuid.Parse(productId); err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			log.Printf("Invalid ID: %v\n", err)
			return
		}

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			http.Error(w, "Invalid price", http.StatusBadRequest)
			log.Printf("Error converting price: %v\n", err)
			return
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			http.Error(w, "Invalid quantity", http.StatusBadRequest)
			log.Printf("Error converting quantity: %v\n", err)
			return
		}

		updatedProduct, err := models.UpdateProduct(productId, name, description, convertedPrice, convertedQuantity)
		if err != nil {
			http.Error(w, "Error updating product", http.StatusInternalServerError)
			log.Printf("Error updating product: %v\n", err)
			return
		}

		log.Printf("Product successfully updated: %v\n", updatedProduct)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
