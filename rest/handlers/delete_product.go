package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	_, err := strconv.Atoi(productID)
	if err != nil {
		util.SendData(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	database.Delete(productID)
	util.SendData(w, "Successfully Deleted Product.", 200)
}
