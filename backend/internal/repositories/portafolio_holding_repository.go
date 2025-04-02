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

func (r *PortafolioHoldingRepository) GetAllByPortafolioID(portafolioID uuid.UUID) ([]models.PortafolioHolding, error) {
	var portafolioHoldings []models.PortafolioHolding
	err := r.db.Preload("Company").Where("portafolio_id = ?", portafolioID).Find(&portafolioHoldings).Error
	return portafolioHoldings, err
}

func (r *PortafolioHoldingRepository) Update(portafolioHolding *models.PortafolioHolding) error {
	return r.db.Save(portafolioHolding).Error
}

func (r *PortafolioHoldingRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.PortafolioHolding{}, id).Error
}
