package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/utils"
	"errors"
)

type CompanyService interface {
	CreateCompanyByTicker(ticker string, name string) (models.Company, error)
	GetCompanyDescription(ticker string, company string) (string, error)
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

func (service *companyService) GetCompanyDescription(ticker string, company string) (string, error) {
	groqClient := utils.NewGroqClient()
	description, err := groqClient.GetCompanySummary(company, ticker)
	if err != nil {
		return "", err
	}
	return description, nil
}
