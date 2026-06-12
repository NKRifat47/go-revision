package product

import "ecommerce/domain"

func (svc *service) Create(p domain.Product) (*domain.Product, error) {
	return svc.prdctRepo.Create(p)
}