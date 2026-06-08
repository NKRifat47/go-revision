package repo

type Product struct {
	ID    int
	Name  string
	Price float64
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ( []*Product, error)
	Delete(productID int) error
	Update(product Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

// Constructor
func NewProductRepo() ProductRepo {
	repo := &productRepo{}

	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil

}
func (r *productRepo) Get(productID int) (*Product, error) {
		for _, product := range r.productList {
		if product.ID == productID {
			return product, nil
		}
	}
	return nil, nil

}
func (r *productRepo) List() ( []*Product, error){
	return r.productList, nil
}
func (r *productRepo) Delete(productID int) error {
	var tempList []*Product
	for _, p := range r.productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}
	r.productList = tempList

	return nil

}
func (r *productRepo) Update(product Product) (*Product, error) {
		for idx, p := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &product
		}
	}

	return &product, nil

}

func generateInitialProducts(r *productRepo) {
	prd1 := &Product{
		ID:    1,
		Name:  "Product 1",
		Price: 10.0,
	}
	prd2 := &Product{
		ID:    2,
		Name:  "Product 2",
		Price: 20.0,
	}

	r.productList = append(r.productList, prd1, prd2)
}