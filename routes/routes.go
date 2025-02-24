package routes

import (
	"net/http"
	"orders-api/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetOrders(w, r)
		} else if r.Method == http.MethodPost {
			handlers.CreateOrder(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
