package repository

import "challenge-10/models"

type ProductRepository interface {
	FindById(id string) *models.Product
	FindAll() []*models.Product
}
