package handlers

import (
	"encoding/json"
	"net/http"
	"orders-api/models"
	"sync"
)

var (
	orders = make(map[int]models.Order)
	mutex  = &sync.Mutex{}
	nextID = 1
)

func GetOrders(w http.ResponseWriter, _ *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	orderList := make([]models.Order, 0, len(orders))
	for _, order := range orders {
		orderList = append(orderList, order)
	}

	json.NewEncoder(w).Encode(orderList)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	order.ID = nextID
	orders[nextID] = order
	nextID++
	mutex.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
