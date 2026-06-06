package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Please give me valid json")
		return
	}

	product := database.Store(newProduct)
	util.SendData(w, product, http.StatusCreated)
}
