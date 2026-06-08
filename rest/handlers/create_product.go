package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/base64"
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

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}