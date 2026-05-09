package database

type Product struct {
	ID    string
	Name  string
	Price float64
}

var Productlist []Product

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

	Productlist = append(Productlist, prd1, prd2)
}
