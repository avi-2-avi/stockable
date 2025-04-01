package repositories

import (
	"backend/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PortafolioHoldingRepository struct {
	db *gorm.DB
}

func NewPortafolioHoldingRepository(db *gorm.DB) *PortafolioHoldingRepository {
	return &PortafolioHoldingRepository{db: db}
}

func (r *PortafolioHoldingRepository) Create(portafolioHolding *models.PortafolioHolding) error {
	return r.db.Create(portafolioHolding).Error
}

func (r *PortafolioHoldingRepository) GetByID(id uuid.UUID) (*models.PortafolioHolding, error) {
	var portafolioHolding models.PortafolioHolding
	err := r.db.First(&portafolioHolding, id).Error
	return &portafolioHolding, err
}

func (r *PortafolioHoldingRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.PortafolioHolding{}, id).Error
}
