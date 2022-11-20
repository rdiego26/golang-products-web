package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"golang-products-web/routes"
	"net/http"
)

const PORT = 8000

func main() {
	routes.RegisterRoutes()
	fmt.Printf("Running server on %d\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
