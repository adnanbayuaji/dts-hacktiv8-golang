package repository

import "challenge-9/models"

type ProductRepository interface {
	FindById(id string) *models.Product
	FindAll() []*models.Product
}
