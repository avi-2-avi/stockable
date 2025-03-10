package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
)

type DataSourceService interface {
	GetAll() ([]models.DataSource, error)
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
