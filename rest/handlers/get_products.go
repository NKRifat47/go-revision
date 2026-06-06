package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := database.List()
	util.SendData(w, products, http.StatusOK)
}
