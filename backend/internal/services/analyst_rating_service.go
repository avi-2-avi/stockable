package services

import (
	"backend/internal/dtos"
	"backend/internal/models"
	"backend/internal/repositories"
)

type AnalystRatingService interface {
	SaveAnalystRating(rating *models.AnalystRating) error
	SaveAnalystRatingsBatch(ratings []models.AnalystRating) error
	GetAll(sortOrder, sortBy, sourceID string, filters map[string]string, page, limit int) ([]models.AnalystRating, int64, error)
	GetIndicators(sourceID string) (dtos.AnalystRatingIndicatorsDTO, error)
	GetMinMaxCPI() (float64, float64, error)
}

type analystRatingService struct {
	ratingRepo *repositories.AnalystRatingRepository
}

func NewAnalystRatingService(ratingRepo *repositories.AnalystRatingRepository) AnalystRatingService {
	return &analystRatingService{
		ratingRepo: ratingRepo,
	}
}

func (service *analystRatingService) SaveAnalystRating(rating *models.AnalystRating) error {
	return service.ratingRepo.Create(rating)
}

func (service *analystRatingService) SaveAnalystRatingsBatch(ratings []models.AnalystRating) error {
	return service.ratingRepo.CreateBatch(ratings)
}

func (service *analystRatingService) GetAll(sortOrder, sortBy, sourceID string, filters map[string]string, page, limit int) ([]models.AnalystRating, int64, error) {
	return service.ratingRepo.GetAll(sortOrder, sortBy, sourceID, filters, page, limit)
}

func (service *analystRatingService) GetMinMaxCPI() (float64, float64, error) {
	return service.ratingRepo.GetMinMaxCPI()
}

func (service *analystRatingService) GetIndicators(sourceID string) (dtos.AnalystRatingIndicatorsDTO, error) {
	return service.ratingRepo.GetIndicators(sourceID)
}
