package services

import (
	"backendAf/config"
	"backendAf/models"
	"backendAf/repositories"
	"backendAf/utils"
	"errors"
)

type userService struct{}

var UserService = userService{}

// SignUp: validate, hash password, persist, return user (without password)
func (s userService) SignUp(input models.User) (*models.User, error) {
	if input.Email == "" || input.Password == "" {
		return nil, errors.New("email and password are required")
	}

	// check exists
	if existing, _ := repositories.UserRepo.FindByEmail(input.Email); existing != nil {
		return nil, errors.New("email already registered")
	}

	hashed, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	input.Password = hashed
	if input.Role == "" {
		input.Role = "user"
	}

	if err := repositories.UserRepo.Create(&input); err != nil {
		return nil, err
	}

	// hide password before returning
	input.Password = ""
	return &input, nil
}

// Login: validate credentials, return token
func (s userService) Login(email, password string) (string, *models.User, error) {
	u, err := repositories.UserRepo.FindByEmail(email)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}
	if err := utils.CheckPassword(u.Password, password); err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	// create JWT
	token, err := config.CreateToken(u.ID, u.Role)
	if err != nil {
		return "", nil, err
	}

	// hide password for return
	u.Password = ""
	return token, u, nil
}
