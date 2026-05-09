package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Go server is Running..."))
	}))

	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handlers.GetProducts),
	))

	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(handlers.CreateProduct),
	))

	mux.Handle("GET /products/{productid}", manager.With(
		http.HandlerFunc(handlers.GetProductByID),
	))
}
