package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
)

type AnalystRatingsService interface {
	SaveAnalystRating(rating *models.AnalystRating) error
	SaveAnalystRatingsBatch(ratings []models.AnalystRating) error
	GetAll() ([]models.AnalystRating, error)
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

func (service *analystRatingsService) GetAll() ([]models.AnalystRating, error) {
	return service.ratingRepo.GetAll()
}
