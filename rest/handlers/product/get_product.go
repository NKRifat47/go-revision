package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")

	_, err := strconv.Atoi(productID)
	if err != nil {
		util.SendData(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product := database.Get(productID)
	if product == nil {
		util.SendError(w, http.StatusNotFound, "Product not found")
		return
	}

	util.SendData(w, product, http.StatusOK)
}
