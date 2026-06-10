package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbcon, err := db.NewConnection()
	if err !=nil{
		fmt.Println(err)
		os.Exit(1)
	}

	productRepo := repo.NewProductRepo()
	userRepo:= repo.NewUserRepo()

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(
		cnf, 
		productHandler, 
		userHandler, 
	)
	server.Start()
}
