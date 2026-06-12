package product

func (svc *service) Delete(productID int) error {
	return svc.prdctRepo.Delete(productID)
}
