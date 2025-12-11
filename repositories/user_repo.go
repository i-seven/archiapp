package repositories

import (
	"backendAf/config"
	"backendAf/models"
)

var UserRepo = userRepository{}

type userRepository struct{}

func (r userRepository) Create(u *models.User) error {
	return config.DB.Create(u).Error
}

func (r userRepository) FindByEmail(email string) (*models.User, error) {
	var u models.User
	if err := config.DB.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r userRepository) FindByID(id uint) (*models.User, error) {
	var u models.User
	if err := config.DB.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
