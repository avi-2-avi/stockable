package repositories

import (
	"data-loader/internal/models"

	"gorm.io/gorm"
)

type AnalystRatingsRepository struct {
	db *gorm.DB
}

func NewAnalystRatingsRepository(db *gorm.DB) *AnalystRatingsRepository {
	return &AnalystRatingsRepository{db: db}
}

func (r *AnalystRatingsRepository) Create(rating *models.AnalystRating) error {
	return r.db.Create(rating).Error
}

func (r *AnalystRatingsRepository) GetByID(id uint) (*models.AnalystRating, error) {
	var rating models.AnalystRating
	err := r.db.First(&rating, id).Error
	if err != nil {
		return nil, err
	}
	return &rating, nil
}
