package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(fullName, email, password string) error
	Login(email, password string) (*models.User, error)
}

type authService struct {
	authRepo *repositories.AuthRepository
}

func NewAuthService(authRepo *repositories.AuthRepository) AuthService {
	return &authService{
		authRepo: authRepo,
	}
}

func (service *authService) Register(fullName, email, password string) error {
	_, err := service.authRepo.GetUserByEmail(email)
	if err == nil {
		return fmt.Errorf("user with email %s already exists", email)
	}

	user := &models.User{
		FullName: fullName,
		Email:    email,
		Password: password,
	}

	if err := user.HashPassword(); err != nil {
		return err
	}

	return service.authRepo.CreateUser(user)
}

func (service *authService) Login(email, password string) (*models.User, error) {
	user, err := service.authRepo.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("user with email %s not found", email)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	return user, nil
}
