package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/util"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	
	header:= r.Header.Get("Authorization")

	if header == "" {
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	headerArr := strings.Split(header, " ")

	if len(headerArr) != 2 {
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	
	accessToken := headerArr[1]

	tokenParts := strings.Split(accessToken, ".")

	if len(tokenParts) != 3 {
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	jwtHeader := tokenParts[0]
	jwtPayload := tokenParts[1]
	jwtSignature := tokenParts[2]

	message := jwtHeader + "." + jwtPayload

	cnf := config.GetConfig()

	byteArrSecret := []byte(cnf.JwtSecretKey)
	byteArrMessage := []byte(message)
	
	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)

	hash := h.Sum(nil)
	newSigneture := base64UrlEncode(hash)

	if newSigneture != jwtSignature{
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

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