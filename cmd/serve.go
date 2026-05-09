package cmd

import (
	"ecommerce/global_router"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	manager.Use(middleware.Logger)
	mux := http.NewServeMux()

	initRoutes(mux, manager)

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on http://localhost:8080")

	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println(err)
	}
}
