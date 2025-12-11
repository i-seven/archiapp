package services

import (
	"backendAf/models"
	"backendAf/repositories"
)

var ProductService = productService{}

type productService struct{}

func (s productService) Create(input models.Product) (models.Product, error) {
	return repositories.ProductRepo.Create(input)
}

func (s productService) GetAll() ([]models.Product, error) {
	return repositories.ProductRepo.GetAll()
}

func (s productService) Get(id string) (models.Product, error) {
	return repositories.ProductRepo.Get(id)
}

func (s productService) Update(id string, input models.Product) (models.Product, error) {
	return repositories.ProductRepo.Update(id, input)
}

func (s productService) Delete(id string) error {
	return repositories.ProductRepo.Delete(id)
}
