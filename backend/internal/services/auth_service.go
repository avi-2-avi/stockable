package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(fullName, email, password, roleName string) (*models.User, error)
	Login(email, password string) (*models.User, error)
	Delete(id uuid.UUID) error
	Update(user *models.User) error
	List() ([]models.User, error)
}

type authService struct {
	authRepo *repositories.AuthRepository
}

func NewAuthService(authRepo *repositories.AuthRepository) AuthService {
	return &authService{
		authRepo: authRepo,
	}
}

func (service *authService) Register(fullName, email, password, roleName string) (*models.User, error) {
	_, err := service.authRepo.GetUserByEmail(email)
	if err == nil {
		return nil, fmt.Errorf("user with email %s already exists", email)
	}

	role, err := service.authRepo.GetRoleByName(roleName)
	if err != nil {
		return nil, fmt.Errorf("role %s not found", roleName)
	}

	user := &models.User{
		FullName: fullName,
		Email:    email,
		Password: password,
		RoleID:   role.ID,
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	err = service.authRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	user, err = service.authRepo.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("user with email %s not found", email)
	}

	return user, nil
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

func (service *authService) Delete(id uuid.UUID) error {
	return service.authRepo.DeleteUser(id)
}

func (service *authService) Update(user *models.User) error {
	return service.authRepo.UpdateUser(user)
}

func (service *authService) List() ([]models.User, error) {
	return service.authRepo.ListUsers()
}
