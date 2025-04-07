package services

import (
	"backend/internal/dtos"
	"backend/internal/models"
	"backend/internal/repositories"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(fullName, email, password, roleName string) (dtos.RegisterUserDTO, error)
	Login(email, password string) (dtos.LoginUserDTO, error)
	Delete(id uuid.UUID) error
	Update(user *models.User) error
	List() ([]dtos.LoginUserDTO, error)
}

type authService struct {
	authRepo *repositories.AuthRepository
}

func NewAuthService(authRepo *repositories.AuthRepository) AuthService {
	return &authService{
		authRepo: authRepo,
	}
}

func (service *authService) Register(fullName, email, password, roleName string) (dtos.RegisterUserDTO, error) {
	_, err := service.authRepo.GetUserByEmail(email)
	if err == nil {
		return dtos.RegisterUserDTO{}, fmt.Errorf("user with email %s already exists", email)
	}

	role, err := service.authRepo.GetRoleByName(roleName)
	if err != nil {
		return dtos.RegisterUserDTO{}, fmt.Errorf("role %s not found", roleName)
	}

	user := &models.User{
		FullName: fullName,
		Email:    email,
		Password: password,
		RoleID:   role.ID,
	}

	if err := user.HashPassword(); err != nil {
		return dtos.RegisterUserDTO{}, err
	}

	err = service.authRepo.CreateUser(user)
	if err != nil {
		return dtos.RegisterUserDTO{}, err
	}

	user, err = service.authRepo.GetUserByEmail(email)
	if err != nil {
		return dtos.RegisterUserDTO{}, fmt.Errorf("user with email %s not found", email)
	}

	return dtos.RegisterUserDTO{
		ID:       user.ID,
		Email:    user.Email,
		FullName: user.FullName,
	}, nil
}

func (service *authService) Login(email, password string) (dtos.LoginUserDTO, error) {
	user, err := service.authRepo.GetUserByEmail(email)
	if err != nil {
		return dtos.LoginUserDTO{}, fmt.Errorf("user with email %s not found", email)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return dtos.LoginUserDTO{}, fmt.Errorf("invalid password")
	}

	return dtos.LoginUserDTO{
		ID:       user.ID,
		Email:    user.Email,
		FullName: user.FullName,
		RoleID:   user.RoleID,
	}, nil
}

func (service *authService) Delete(id uuid.UUID) error {
	return service.authRepo.DeleteUser(id)
}

func (service *authService) Update(user *models.User) error {
	return service.authRepo.UpdateUser(user)
}

func (service *authService) List() ([]dtos.LoginUserDTO, error) {
	users, err := service.authRepo.ListUsers()
	if err != nil {
		return nil, err
	}

	var userDtos []dtos.LoginUserDTO
	for _, user := range users {
		userDtos = append(userDtos, dtos.LoginUserDTO{
			ID:       user.ID,
			Email:    user.Email,
			FullName: user.FullName,
			RoleID:   user.RoleID,
		})
	}

	return userDtos, nil
}
