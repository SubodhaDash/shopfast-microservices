package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Order struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateOrderRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type HealthResponse struct {
	Status    string `json:"status"`
	Service   string `json:"service"`
	Timestamp string `json:"timestamp"`
}

var orders []Order
var orderCounter = 1

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HealthResponse{
		Status:    "healthy",
		Service:   "order-service",
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

func getOrdersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	order := Order{
		ID:        orderCounter,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Status:    "pending",
		CreatedAt: time.Now(),
	}
	orders = append(orders, order)
	orderCounter++

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getOrdersHandler(w, r)
		case http.MethodPost:
			createOrderHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Printf("Order Service running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}