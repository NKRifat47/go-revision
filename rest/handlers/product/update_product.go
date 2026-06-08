package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

type ReqUpdateProduct struct {
	Name  string
	Price float64
}


func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var req ReqUpdateProduct

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Please give me valid json")
		return
	}

	_, err = h.productRepo.Update(repo.Product{
		ID: pId,
		Name: req.Name,
		Price: req.Price,
	})
	if err !=nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	util.SendData(w, http.StatusOK, "Product Updated")
}
