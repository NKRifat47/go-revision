package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handlers.GetProducts),
		middleware.Arekta,
	))

	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))

	mux.Handle("GET /products/{productid}", http.HandlerFunc(handlers.GetProductByID))
}
