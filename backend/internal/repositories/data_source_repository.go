package repositories

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

type DataSourceRepository struct {
	db *gorm.DB
}

func NewDataSourceRepository(db *gorm.DB) *DataSourceRepository {
	return &DataSourceRepository{db: db}
}

func (r *DataSourceRepository) Create(dataSource *models.DataSource) error {
	return r.db.Create(dataSource).Error
}

func (r *DataSourceRepository) GetByID(id uint) (*models.DataSource, error) {
	var dataSource models.DataSource
	err := r.db.First(&dataSource, id).Error
	return &dataSource, err
}

func (r *DataSourceRepository) GetByName(name string) (*models.DataSource, error) {
	var dataSource models.DataSource
	err := r.db.Where("name = ?", name).First(&dataSource).Error
	return &dataSource, err
}

func (r *DataSourceRepository) Delete(id uint) error {
	return r.db.Delete(&models.DataSource{}, id).Error
}
