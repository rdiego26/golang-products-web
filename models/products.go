package models

import "golang-products-web/database"

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
