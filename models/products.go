package models

import (
	"github.com/google/uuid"
	"golang-products-web/database"
	"log"
)

type Product struct {
	Id          string
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {
	dbConnection := database.GetConnection()

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
	defer dbConnection.Close()

	return products
}

func CreateProduct(name, description string, price float64, quantity int) {
	db := database.GetConnection()

	log.Println("Received:", name, description, price, quantity)

	insertData, err := db.Prepare("INSERT INTO products(name, description, price, quantity) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	_, err = insertData.Exec(name, description, price, quantity)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func DeleteProduct(productId string) {
	db := database.GetConnection()

	removeData, err := db.Prepare("DELETE FROM products WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	_, err = removeData.Exec(productId)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func GetProduct(productId uuid.UUID) Product {
	dbConnection := database.GetConnection()

	fetchedProduct, err := dbConnection.Query("SELECT * FROM products WHERE id=$1", productId)
	if err != nil {
		panic(err.Error())
	}

	defer dbConnection.Close()

	result := Product{}

	for fetchedProduct.Next() {
		var id, name, description string
		var price float64
		var quantity int

		err = fetchedProduct.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		result.Id = id
		result.Name = name
		result.Description = description
		result.Price = price
		result.Quantity = quantity
	}

	return result
}

func UpdateProduct(productId string, name string, description string, price float64, quantity int) Product {
	db := database.GetConnection()

	updateData, err := db.Prepare("UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	parsedProductId, err := uuid.Parse(productId)
	if err != nil {
		panic(err.Error())
	}

	_, err = updateData.Exec(name, description, price, quantity, parsedProductId)
	if err != nil {
		panic(err.Error())
	}

	return GetProduct(parsedProductId)
}
