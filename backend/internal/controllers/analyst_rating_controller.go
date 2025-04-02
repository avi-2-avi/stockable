package controllers

import (
	"backend/internal/services"
	"backend/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AnalystRatingController struct {
	RatingService services.AnalystRatingService
}

func NewAnalystRatingController(ratingService services.AnalystRatingService) *AnalystRatingController {
	return &AnalystRatingController{
		RatingService: ratingService,
	}
}

func (controller *AnalystRatingController) GetRatings(context *gin.Context) {
	sortBy, sortOrder, sourceID, pageNumber, limitNumber, filters := parseQueryParams(context)

	ratingsDTOs, total, err := controller.RatingService.GetAll(sortOrder, sortBy, sourceID, filters, pageNumber, limitNumber)
	if err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get ratings",
		})
		return
	}

	utils.Respond(context, utils.APIResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Body: gin.H{
			"page":    pageNumber,
			"limit":   limitNumber,
			"total":   total,
			"ratings": ratingsDTOs,
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

func (controller *AnalystRatingController) GetRatingsIndicators(context *gin.Context) {
	sourceID := context.Query("source_id")
	indicators, err := controller.RatingService.GetIndicators(sourceID)
	if err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get indicators",
		})
		return
	}

	utils.Respond(context, utils.APIResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Body: gin.H{
			"indicators": indicators,
		},
	})
}

func (controller *AnalystRatingController) GetDashboardRatings(context *gin.Context) {
	sourceID := context.Query("source_id")
	dashboardData, err := controller.RatingService.GetDashboardRatings(sourceID)
	if err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get ratings",
		})
		return
	}

	utils.Respond(context, utils.APIResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Body: gin.H{
			"data": dashboardData,
		},
	})

}
