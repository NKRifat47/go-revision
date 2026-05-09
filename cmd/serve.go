package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Cors,
		middleware.Preflight,
		middleware.Logger,
	)

	mux := http.NewServeMux()

	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	fmt.Println("Server is running on http://localhost:8080")

	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		fmt.Println(err)
	}
}
