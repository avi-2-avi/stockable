package adapters

import "data-loader/internal/models"

type RatingAdapter interface {
	FetchData() ([]models.AnalystRating, error)
}
