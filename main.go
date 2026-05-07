package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am Rifat.")
}

type Product struct {
	ID    string
	Name  string
	Price float64
}

var products []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	handelCors(w)
	handlePreflightRequest(w, r)

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sendData(w, products, 200)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	handelCors(w)
	handlePreflightRequest(w, r)

	if r.Method == http.MethodPost {
		var newProduct Product
		err := json.NewDecoder(r.Body).Decode(&newProduct)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		newProduct.ID = strconv.Itoa(len(products) + 1)
		products = append(products, newProduct)
		sendData(w, newProduct, 201)

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handelCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func handlePreflightRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
	}
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)

	mux.HandleFunc("/products", getProducts)
	mux.HandleFunc("/createProduct", createProduct)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", mux)
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
