package product

import "ecommerce/domain"

func (svc *service) Update(product domain.Product) (*domain.Product, error) {
	return svc.prdctRepo.Update(product)
}