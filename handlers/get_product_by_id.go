package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("productid")

	_, err := strconv.Atoi(productID)
	if err != nil {
		util.SendData(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for _, product := range database.Productlist {
		if product.ID == productID {
			util.SendData(w, product, http.StatusOK)
			return
		}
	}

	util.SendData(w, "Product not found", http.StatusNotFound)
}
