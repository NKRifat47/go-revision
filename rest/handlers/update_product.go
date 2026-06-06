package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	_, err := strconv.Atoi(productID)
	if err != nil {
		util.SendData(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Please give me valid json")
		return
	}

	newProduct.ID = productID

	database.Update(newProduct)

	util.SendData(w, "Successfully Updated Product.", 200)
}
