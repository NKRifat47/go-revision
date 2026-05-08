package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

var products []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	sendData(w, products, 200)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	newProduct.ID = strconv.Itoa(len(products) + 1)
	products = append(products, newProduct)
	sendData(w, newProduct, 201)
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(getProducts))
	mux.Handle("POST /createProduct", http.HandlerFunc(createProduct))

	fmt.Println("Server is running on port 8080")

	globalRouter := globalRouter(mux)
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println(err)
	}
}

func init() {
	prd1 := Product{
		ID:    "1",
		Name:  "Product 1",
		Price: 10.0,
	}
	prd2 := Product{
		ID:    "2",
		Name:  "Product 2",
		Price: 20.0,
	}

	products = append(products, prd1, prd2)
}

func globalRouter(mux *http.ServeMux) http.Handler {
	handleAllRequests := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		mux.ServeHTTP(w, r)
	}

	handler := http.HandlerFunc(handleAllRequests)

	return handler

}
