package adapters

import "backend/internal/models"

type RatingAdapter interface {
	FetchData() ([]models.AnalystRating, error)
}
