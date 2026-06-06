package database

import "strconv"

var productList []Product

type Product struct {
	ID    string
	Name  string
	Price float64
}

func Store(p Product) Product {
	p.ID = strconv.Itoa(len(productList) + 1)
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(productID string) *Product {
	for _, product := range productList {
		if product.ID == productID {
			return &product
		}
	}
	return nil
}

func Update(product Product) {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product
			return
		}
	}
}

func Delete(productID string) {
	var tempList []Product
	for _, p := range productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}
	productList = tempList
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

	productList = append(productList, prd1, prd2)
}
