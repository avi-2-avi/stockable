package services

import (
	"backend/internal/models"
	"backend/internal/repositories"

	"github.com/google/uuid"
)

type DataSourceService interface {
	GetAll() ([]models.DataSource, error)
	Update(id string, isVisible bool) (*models.DataSource, error)
}

type dataSourceService struct {
	sourceRepo *repositories.DataSourceRepository
}

func NewDataSourceService(sourceRepo *repositories.DataSourceRepository) DataSourceService {
	return &dataSourceService{
		sourceRepo: sourceRepo,
	}
}

func (service *dataSourceService) GetAll() ([]models.DataSource, error) {
	return service.sourceRepo.GetAll()
}

func (service *dataSourceService) Update(id string, isVisible bool) (*models.DataSource, error) {
	return service.sourceRepo.Update(uuid.MustParse(id), isVisible)
}
