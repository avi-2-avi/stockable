package controllers

import (
	"backend/internal/dtos"
	"backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AnalystRatingController struct {
	RatingService services.AnalystRatingsService
}

func NewAnalystRatingController(ratingService services.AnalystRatingsService) *AnalystRatingController {
	return &AnalystRatingController{
		RatingService: ratingService,
	}
}

func (controller *AnalystRatingController) GetRatings(context *gin.Context) {
	ratings, err := controller.RatingService.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get ratings",
		})
		return
	}

	var ratingDTOs []*dtos.AnalystRatingDTO
	for _, rating := range ratings {
		ratingDTOs = append(ratingDTOs, &dtos.AnalystRatingDTO{
			ID:         rating.ID,
			Ticker:     rating.Ticker,
			TargetFrom: rating.TargetFrom,
			TargetTo:   rating.TargetTo,
			Company:    rating.Company,
			Action:     rating.Action,
			Brokerage:  rating.Brokerage,
			RatingFrom: rating.RatingFrom,
			RatingTo:   rating.RatingTo,
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"ratings": ratingDTOs,
	})
}
