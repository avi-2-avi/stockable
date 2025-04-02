package services

import (
	"backend/internal/models"
	"backend/internal/repositories"

	"github.com/google/uuid"
)

type PortfolioService interface {
	GetPortafoliosByUserID(userID uuid.UUID) ([]models.Portafolio, error)
	UpdatePortafolio(portfolio *models.Portafolio) error
	DeletePortafolio(id uuid.UUID) error
	CreatePortafolio(portfolio *models.Portafolio) error
}

type portafolioService struct {
	portafolioRepo *repositories.PortafolioRepository
}

func NewPortafolioService(portafolioRepo *repositories.PortafolioRepository) PortfolioService {
	return &portafolioService{
		portafolioRepo: portafolioRepo,
	}
}

func (service *portafolioService) GetPortafoliosByUserID(userID uuid.UUID) ([]models.Portafolio, error) {
	return service.portafolioRepo.GetAllByUserID(userID)
}

func (service *portafolioService) UpdatePortafolio(portfolio *models.Portafolio) error {
	return service.portafolioRepo.Update(portfolio)
}

func (service *portafolioService) DeletePortafolio(id uuid.UUID) error {
	return service.portafolioRepo.Delete(id)
}

func (service *portafolioService) CreatePortafolio(portfolio *models.Portafolio) error {
	return service.portafolioRepo.Create(portfolio)
}
