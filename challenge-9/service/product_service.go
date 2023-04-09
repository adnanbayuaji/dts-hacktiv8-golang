package service

import (
	"errors"

	"challenge-9/models"
	"challenge-9/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id string) (*models.Product, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}

func (service ProductService) GetAllProduct() ([]*models.Product, error) {
	products := service.Repository.FindAll()
	if len(products) == 0 {
		return nil, errors.New("all products not found")
	}

	return products, nil
}
