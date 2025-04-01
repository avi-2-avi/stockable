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

func (r *PortafolioRepository) GetByID(id uuid.UUID) (*models.Portafolio, error) {
	var portafolio models.Portafolio
	err := r.db.First(&portafolio, id).Error
	return &portafolio, err
}

func (r *PortafolioRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Portafolio{}, id).Error
}
