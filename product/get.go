package product

import "ecommerce/domain"

func (svc *service) Get(productID int) (*domain.Product, error) {
	return svc.prdctRepo.Get(productID)
}
