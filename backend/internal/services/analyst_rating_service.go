package services

import (
	"backend/internal/dtos"
	"backend/internal/models"
	"backend/internal/repositories"
)

type AnalystRatingsService interface {
	SaveAnalystRating(rating *models.AnalystRating) error
	SaveAnalystRatingsBatch(ratings []models.AnalystRating) error
	GetAll(sortOrder, sortBy, sourceID string, filters map[string]string, page, limit int) ([]models.AnalystRating, int64, error)
	GetIndicators(sourceID string) (dtos.AnalystRatingIndicatorsDTO, error)
	GetRecommendations() ([]models.AnalystRating, error)
}

type analystRatingsService struct {
	ratingRepo *repositories.AnalystRatingsRepository
}

func NewAnalystRatingsService(ratingRepo *repositories.AnalystRatingsRepository) AnalystRatingsService {
	return &analystRatingsService{
		ratingRepo: ratingRepo,
	}
}

func (service *analystRatingsService) SaveAnalystRating(rating *models.AnalystRating) error {
	return service.ratingRepo.Create(rating)
}

func (service *analystRatingsService) SaveAnalystRatingsBatch(ratings []models.AnalystRating) error {
	return service.ratingRepo.CreateBatch(ratings)
}

func (service *analystRatingsService) GetAll(sortOrder, sortBy, sourceID string, filters map[string]string, page, limit int) ([]models.AnalystRating, int64, error) {
	return service.ratingRepo.GetAll(sortOrder, sortBy, sourceID, filters, page, limit)
}

func (service *analystRatingsService) GetIndicators(sourceID string) (dtos.AnalystRatingIndicatorsDTO, error) {
	return service.ratingRepo.GetIndicators(sourceID)
}

// TODO: Implement the GetRecommendations method

func (service *analystRatingsService) GetRecommendations() ([]models.AnalystRating, error) {
	return service.ratingRepo.GetRecommendations()
}
