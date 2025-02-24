package main

import (
	"net/http"
	"orders-api/routes"
)

func main() {
	routes.RegisterRoutes()
	http.ListenAndServe(":8080", nil)
}
