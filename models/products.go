package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang-products-web/database"
	"log"
)

// Product structure
type Product struct {
	Id          string
	Name        string
	Description string
	Price       float64
	Quantity    int
}

// Get all products
func GetAllProducts() ([]Product, error) {
	db := database.GetConnection()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, description, price, quantity FROM products ORDER BY name")
	if err != nil {
		return nil, fmt.Errorf("error fetching products: %w", err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity)
		if err != nil {
			return nil, fmt.Errorf("error scanning product: %w", err)
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over products: %w", err)
	}

	return products, nil
}

// Create a product
func CreateProduct(name, description string, price float64, quantity int) error {
	if name == "" || price < 0 || quantity < 0 {
		return errors.New("invalid values for product creation")
	}

	db := database.GetConnection()
	defer db.Close()

	query := `INSERT INTO products(name, description, price, quantity) VALUES($1, $2, $3, $4)`
	_, err := db.Exec(query, name, description, price, quantity)
	if err != nil {
		return fmt.Errorf("error inserting product: %w", err)
	}

	log.Printf("Product created: %s", name)
	return nil
}

// Delete a product
func DeleteProduct(productId string) error {
	if _, err := uuid.Parse(productId); err != nil {
		return fmt.Errorf("invalid UUID: %w", err)
	}

	db := database.GetConnection()
	defer db.Close()

	query := `DELETE FROM products WHERE id=$1`
	result, err := db.Exec(query, productId)
	if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no product found to delete with id: %s", productId)
	}

	log.Printf("Product successfully deleted: %s", productId)
	return nil
}

// Get a specific product
func GetProduct(productId string) (*Product, error) {
	if _, err := uuid.Parse(productId); err != nil {
		return nil, fmt.Errorf("invalid UUID: %w", err)
	}

	db := database.GetConnection()
	defer db.Close()

	query := `SELECT id, name, description, price, quantity FROM products WHERE id=$1`
	row := db.QueryRow(query, productId)

	var product Product
	err := row.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("product not found with id: %s", productId)
	}
	if err != nil {
		return nil, fmt.Errorf("error fetching product: %w", err)
	}

	return &product, nil
}

// Update a product
func UpdateProduct(productId, name, description string, price float64, quantity int) (*Product, error) {
	if _, err := uuid.Parse(productId); err != nil {
		return nil, fmt.Errorf("invalid UUID: %w", err)
	}

	if name == "" || price < 0 || quantity < 0 {
		return nil, errors.New("invalid values for product update")
	}

	db := database.GetConnection()
	defer db.Close()

	query := `UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5`
	result, err := db.Exec(query, name, description, price, quantity, productId)
	if err != nil {
		return nil, fmt.Errorf("error updating product: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, fmt.Errorf("no product found to update with id: %s", productId)
	}

	// Return the updated product
	return GetProduct(productId)
}
