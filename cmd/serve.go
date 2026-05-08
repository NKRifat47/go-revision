package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /rifat", middleware.Logger(http.HandlerFunc(handlers.Test)))

	mux.Handle("GET /products", middleware.Logger(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /Products", middleware.Logger(http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle("GET /products/{productid}", middleware.Logger(http.HandlerFunc(handlers.GetProductByID)))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on port 8080")

	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println(err)
	}
}
