package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
)

type AnalystRatingsService struct {
	ratingsRepo *repositories.AnalystRatingsRepository
}

func NewAnalystRatingsService(ratingsRepo *repositories.AnalystRatingsRepository) *AnalystRatingsService {
	return &AnalystRatingsService{ratingsRepo: ratingsRepo}
}

func (service *AnalystRatingsService) SaveAnalystRating(rating *models.AnalystRating) error {
	return service.ratingsRepo.Create(rating)
}

func (service *AnalystRatingsService) SaveAnalystRatingsBatch(ratings []models.AnalystRating) error {
	return service.ratingsRepo.CreateBatch(ratings)
}
