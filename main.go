package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
)

type Product struct {
	Id          string
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
	dbConnection := createDatabaseConnection()

	allProducts, err := dbConnection.Query("SELECT * FROM products ORDER BY name")
	if err != nil {
		panic(err.Error())
	}

	product := Product{}
	var products []Product

	for allProducts.Next() {
		var id, name, description string
		var price float64
		var quantity int

		err = allProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity

		products = append(products, product)
	}

	temp.ExecuteTemplate(w, "Index", products)
	defer dbConnection.Close()
}

func createDatabaseConnection() *sql.DB {
	connectionStr := "user=postgres dbname=postgres password=postgres host=localhost port=35432 sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err.Error())
	}

	return db
}
