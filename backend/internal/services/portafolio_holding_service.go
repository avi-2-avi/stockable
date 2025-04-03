package services

import (
	"backend/internal/models"
	"backend/internal/repositories"

	"github.com/google/uuid"
)

type PortafolioHoldingService interface {
	GetAllByPortafolioID(portafolioID uuid.UUID) ([]models.PortafolioHolding, error)
	Create(portafolioHolding *models.PortafolioHolding) error
	Update(portafolioHolding *models.PortafolioHolding) error
	Delete(id uuid.UUID) error
}

type portafolioHoldingService struct {
	portafolioHoldingRepo *repositories.PortafolioHoldingRepository
}

func NewPortafolioHoldingService(portafolioHoldingRepo *repositories.PortafolioHoldingRepository) PortafolioHoldingService {
	return &portafolioHoldingService{
		portafolioHoldingRepo: portafolioHoldingRepo,
	}
}

func (s *portafolioHoldingService) GetAllByPortafolioID(portafolioID uuid.UUID) ([]models.PortafolioHolding, error) {
	return s.portafolioHoldingRepo.GetAllByPortafolioID(portafolioID)
}

func (s *portafolioHoldingService) Create(portafolioHolding *models.PortafolioHolding) error {
	return s.portafolioHoldingRepo.Create(portafolioHolding)
}

func (s *portafolioHoldingService) Update(portafolioHolding *models.PortafolioHolding) error {
	return s.portafolioHoldingRepo.Update(portafolioHolding)
}

func (s *portafolioHoldingService) Delete(id uuid.UUID) error {
	return s.portafolioHoldingRepo.Delete(id)
}
