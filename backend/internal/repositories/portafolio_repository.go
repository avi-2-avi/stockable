package repositories

import (
	"backend/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PortafolioRepository struct {
	db *gorm.DB
}

func NewPortafolioRepository(db *gorm.DB) *PortafolioRepository {
	return &PortafolioRepository{db: db}
}

func (r *PortafolioRepository) Create(portafolio *models.Portafolio) error {
	return r.db.Create(portafolio).Error
}

func (r *PortafolioRepository) GetAllByUserID(userID uuid.UUID) ([]models.Portafolio, error) {
	var portafolios []models.Portafolio
	err := r.db.Where("user_id = ?", userID).Preload("DataSource").Find(&portafolios).Error

	return portafolios, err
}

func (r *PortafolioRepository) Update(portafolio *models.Portafolio) error {
	return r.db.Save(portafolio).Error
}

func (r *PortafolioRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Portafolio{}, id).Error
}

func (r *PortafolioRepository) GetByID(id uuid.UUID) (*models.Portafolio, error) {
	var portafolio models.Portafolio
	err := r.db.Preload("DataSource").First(&portafolio, id).Error
	return &portafolio, err
}
