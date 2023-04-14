package repository

import "final-project/models"

type ProductRepository interface {
	FindById(id string) *models.Product
	FindAll() []*models.Product
}
