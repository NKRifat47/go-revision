package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/base64"
	"encoding/json"
	"net/http"
)

type ReqCreateProduct struct {
	Name  string
	Price float64
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	
	var req ReqCreateProduct

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Please give me valid json")
		return
	}

	product, err := h.productRepo.Create(repo.Product{
		Name: req.Name,
		Price: req.Price,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	util.SendData(w, http.StatusCreated, product)
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}