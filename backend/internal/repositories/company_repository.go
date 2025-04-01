package repositories

import (
	"backend/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

func (r *CompanyRepository) FindByTicker(ticker string) (models.Company, error) {
	var company models.Company
	err := r.db.Where("ticker = ?", ticker).First(&company).Error
	if err != nil {
		return company, err
	}
	return company, nil
}

func (r *CompanyRepository) Create(company *models.Company) error {
	return r.db.Create(company).Error
}

func (r *CompanyRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Company{}, id).Error
}
