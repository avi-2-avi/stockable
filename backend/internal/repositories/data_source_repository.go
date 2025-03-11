package repositories

import (
	"backend/internal/models"

	"github.com/google/uuid"
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

func (r *DataSourceRepository) GetAll() ([]models.DataSource, error) {
	var dataSources []models.DataSource
	err := r.db.Find(&dataSources).Error
	return dataSources, err
}

func (r *DataSourceRepository) GetByID(id uuid.UUID) (*models.DataSource, error) {
	var dataSource models.DataSource
	err := r.db.First(&dataSource, id).Error
	return &dataSource, err
}

func (r *DataSourceRepository) GetByName(name string) (*models.DataSource, error) {
	var dataSource models.DataSource
	err := r.db.Where("name = ?", name).First(&dataSource).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &dataSource, nil
}

func (r *DataSourceRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.DataSource{}, id).Error
}
