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
			Code:    http.StatusInternalServerError,
			Message: "Failed to get sources",
		})
		return
	}

	var sourceDTOs []*dtos.DataSourceDTO
	for _, source := range sources {
		sourceDTOs = append(sourceDTOs, &dtos.DataSourceDTO{
			ID:   source.ID,
			Name: source.Name,
		})
	}

	utils.Respond(context, utils.APIResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    sourceDTOs,
	})
}
