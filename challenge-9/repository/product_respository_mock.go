package repository

import (
	"challenge-9/models"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindById(id string) *models.Product {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}

	product := arguments.Get(0).(models.Product)

	return &product
}

func (repository *ProductRepositoryMock) FindAll() []*models.Product {
	arguments := repository.Mock.Called()
	return arguments.Get(0).([]*models.Product)
}
