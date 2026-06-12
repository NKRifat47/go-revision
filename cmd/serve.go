package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/product"
	"ecommerce/repo"
	"ecommerce/rest"
	prodctHandler "ecommerce/rest/handlers/product"
	usrHandler "ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"ecommerce/user"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err !=nil{
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err !=nil{
		fmt.Println(err)
		os.Exit(1)
	}

	//Repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo:= repo.NewUserRepo(dbCon)

	//Domains
	usrSvc := user.NewService(userRepo)
	prdctSvc := product.NewService(productRepo)

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := prodctHandler.NewHandler(middlewares, prdctSvc)
	userHandler := usrHandler.NewHandler(cnf, usrSvc)

	server := rest.NewServer(
		cnf, 
		productHandler, 
		userHandler, 
	)
	server.Start()
}
