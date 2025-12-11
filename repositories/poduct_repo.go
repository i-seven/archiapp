package repositories

import (
	"backendAf/config"
	"backendAf/models"
	"fmt"
	"strconv"
)

var ProductRepo = productRepository{}

type productRepository struct{}

func (r productRepository) Create(p models.Product) (models.Product, error) {
	err := config.DB.Create(&p).Error
	return p, err
}

func (r productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := config.DB.Preload("ArchiStyles").
		Preload("Ceilings").
		Preload("Garages").
		Preload("RoofDetails").
		Preload("SpecialFeatures").
		Find(&products).Error
	return products, err
}

func (r productRepository) Get(id string) (models.Product, error) {
	var p models.Product
	err := config.DB.Preload("ArchiStyles").
		Preload("Ceilings").
		Preload("Garages").
		Preload("RoofDetails").
		Preload("SpecialFeatures").
		First(&p, "pid = ?", id).Error
	return p, err
}

func (r productRepository) Update(id string, input models.Product) (models.Product, error) {

	number, conver_err := strconv.Atoi(id)
	if conver_err != nil {
		fmt.Println("Conversion error:", conver_err)
		return input, conver_err
	}

	input.PID = number
	err := config.DB.Save(&input).Error
	return input, err
}

func (r productRepository) Delete(id string) error {
	return config.DB.Delete(&models.Product{}, "pid = ?", id).Error
}
