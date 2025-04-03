package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"errors"
)

type CompanyService interface {
	CreateCompanyByTicker(ticker string, name string) (models.Company, error)
}

type companyService struct {
	companyRepo *repositories.CompanyRepository
}

func NewCompanyService(companyRepo *repositories.CompanyRepository) CompanyService {
	return &companyService{
		companyRepo: companyRepo,
	}
}

func (service *companyService) CreateCompanyByTicker(ticker string, name string) (models.Company, error) {
	company, err := service.companyRepo.FindByTicker(ticker)
	if err == nil {
		return company, nil
	}

	newCompany := models.Company{
		Ticker: ticker,
		Name:   name,
	}

	err = service.companyRepo.Create(&newCompany)
	if err != nil {
		return models.Company{}, errors.New("failed to create company")
	}

	return newCompany, nil
}
