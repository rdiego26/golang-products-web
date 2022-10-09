package routes

import "net/http"
import "golang-products-web/controllers"

func RegisterRoutes() {
	http.HandleFunc("/", controllers.Index)
}
