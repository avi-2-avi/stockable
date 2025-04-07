package repositories

import (
	"backend/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (r *AuthRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *AuthRepository) CreateRole(role *models.Role) error {
	return r.db.Create(role).Error
}

func (r *AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Role").Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *AuthRepository) GetRoleByName(name string) (*models.Role, error) {
	var role models.Role
	err := r.db.Where("name = ?", name).First(&role).Error
	return &role, err
}

func (r *AuthRepository) DeleteUser(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&models.User{}).Error
}

func (r *AuthRepository) DeleteRole(id uint) error {
	return r.db.Delete(&models.Role{}, id).Error
}

func (r *AuthRepository) UpdateUser(user *models.User) error {
	var existingUser models.User
	if err := r.db.First(&existingUser, "id = ?", user.ID).Error; err != nil {
		return err
	}

	if user.Password != "" && user.Password != existingUser.Password {
		if err := user.HashPassword(); err != nil {
			return err
		}
	}

	return r.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(user).Error
}

func (r *AuthRepository) ListUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}
