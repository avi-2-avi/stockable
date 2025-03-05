package controllers

import (
	"backend/internal/dtos"
	"backend/internal/services"
	"backend/internal/utils"
	"net/http"
	"strconv"

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
	sortBy, sortOrder, sourceID, pageNumber, limitNumber, filters := parseQueryParams(context)

	ratings, err := controller.RatingService.GetAll(sortOrder, sortBy, sourceID, filters, pageNumber, limitNumber)
	if err != nil {
		utils.Respond(context, utils.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get ratings",
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
			RatedAt:    rating.RatedAt,
		})
	}

	utils.Respond(context, utils.APIResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data: gin.H{
			"page":    pageNumber,
			"limit":   limitNumber,
			"total":   len(ratings),
			"ratings": ratingDTOs,
		},
	})
}

func parseQueryParams(context *gin.Context) (string, string, string, int, int, map[string]string) {
	sortBy := context.DefaultQuery("sort_by", "rated_at")
	sortOrder := context.DefaultQuery("sort_order", "asc")
	sourceID := context.Query("source_id")
	page := context.DefaultQuery("page", "1")
	limit := context.DefaultQuery("limit", "10")

	pageNumber := parseInt(page, 1)
	limitNumber := parseInt(limit, 10)

	filters := extractFilters(context)

	return sortBy, sortOrder, sourceID, pageNumber, limitNumber, filters
}

func parseInt(value string, defaultValue int) int {
	number, err := strconv.Atoi(value)
	if err != nil || number < 1 {
		return defaultValue
	}
	return number
}

func extractFilters(context *gin.Context) map[string]string {
	filters := make(map[string]string)

	for key, values := range context.Request.URL.Query() {
		if key == "sort_by" || key == "sort_order" || key == "source_id" || key == "page" || key == "limit" {
			continue
		}

		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	return filters
}

func (controller *AnalystRatingController) GetRecommendations(context *gin.Context) {
	recommendations, err := controller.RatingService.GetRecommendations()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get recommendations",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"recommendations": recommendations,
	})
}
