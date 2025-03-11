package repositories

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

type AdapterLogRepository struct {
	db *gorm.DB
}

func NewAdapterLogRepository(db *gorm.DB) *AdapterLogRepository {
	return &AdapterLogRepository{db: db}
}

func (r *AdapterLogRepository) Create(log *models.AdapterLog) error {
	return r.db.Create(log).Error
}

func (r *AdapterLogRepository) GetByName(adapterName string) (*models.AdapterLog, error) {
	var log models.AdapterLog
	err := r.db.Where("adapter_name = ?", adapterName).First(&log).Error
	return &log, err
}

func (r *AdapterLogRepository) Delete(id uint) error {
	return r.db.Delete(&models.AdapterLog{}, id).Error
}
