package service

import (
	"testing"

	"challenge-9/models"
	"challenge-9/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProduct(t *testing.T) {
	product := models.Product{
		Title:       "Bekas Bantal",
		Description: "Bantal ini tidak laku dijual",
	}

	productRepository.Mock.On("FindById", "2").Return(product)
	result, err := productService.GetOneProduct("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, product.Title, result.Title, "result has to be 'Bekas bantal'")
	assert.Equal(t, product.Description, result.Description, "result has to be 'Bantal ini tidak laku dijual'")
	assert.Equal(t, &product, result, "result has to be a product data with id '2'")
}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindById", "1").Return(nil)
	product, err := productService.GetOneProduct("1")
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}

func TestProductServiceGetAllProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindAll").Return(make([]*models.Product, 0))
	products, err := productService.GetAllProduct()
	assert.Nil(t, products)
	assert.NotNil(t, err)
	assert.Equal(t, "all products not found", err.Error(), "error response has to be 'all products no found'")
}
