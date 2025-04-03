package controllers

import (
	"backend/internal/dtos"
	"backend/internal/services"
	"backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DataSourceController struct {
	SourceService services.DataSourceService
}

func NewDataSourceController(sourceService services.DataSourceService) *DataSourceController {
	return &DataSourceController{
		SourceService: sourceService,
	}
}

func (controller *DataSourceController) GetSources(context *gin.Context) {
	sources, err := controller.SourceService.GetAll()
	if err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get sources",
		})
		return
	}

	var sourceDTOs []*dtos.DataSourceDTO
	for _, source := range sources {
		sourceDTOs = append(sourceDTOs, &dtos.DataSourceDTO{
			ID:        source.ID,
			Name:      source.Name,
			IsVisible: source.IsVisible,
		})
	}

	utils.Respond(context, utils.APIResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Body:    sourceDTOs,
	})
}

type UpdateDataSourceDTO struct {
	IsVisible bool `json:"is_visible"`
}

func (controller *DataSourceController) UpdateSource(context *gin.Context) {
	id := context.Param("id")

	var updateDTO UpdateDataSourceDTO
	if err := context.ShouldBindJSON(&updateDTO); err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update source",
		})
		return
	}

	_, err := controller.SourceService.Update(id, updateDTO.IsVisible)
	if err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to update source",
		})
		return
	}

	utils.Respond(context, utils.APIResponse{
		Status:  http.StatusOK,
		Message: "Success",
	})
}
