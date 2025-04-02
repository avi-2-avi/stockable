package controllers

import (
	"backend/internal/dtos"
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PortafolioController struct {
	PortafolioService services.PortfolioService
}

func NewPortafolioController(portafolioService services.PortfolioService) *PortafolioController {
	return &PortafolioController{
		PortafolioService: portafolioService,
	}
}

func (controller *PortafolioController) GetPortafolios(context *gin.Context) {
	user_id := context.Param("user_id")
	portafolios, err := controller.PortafolioService.GetPortafoliosByUserID(uuid.MustParse(user_id))
	if err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get portafolios",
		})
		return
	}

	var portafolioDTOs []dtos.PortafolioDTO
	for _, portafolio := range portafolios {
		portafolioDTOs = append(portafolioDTOs, dtos.PortafolioDTO{
			ID:       portafolio.ID,
			Name:     portafolio.Name,
			Category: portafolio.Category,
			DataSource: dtos.DataSourceDTO{
				ID:   portafolio.DataSource.ID,
				Name: portafolio.DataSource.Name,
			},
		})
	}

	utils.Respond(context, utils.APIResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Body:    portafolioDTOs,
	})
}

type UpdatePortafolioRequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

func (controller *PortafolioController) UpdatePortafolio(context *gin.Context) {
	id := context.Param("id")

	var updatePortafolioRequest UpdatePortafolioRequest
	if err := context.ShouldBindJSON(&updatePortafolioRequest); err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update portafolio",
		})
		return
	}

	portafolio := models.Portafolio{ID: uuid.MustParse(id)}

	if err := controller.PortafolioService.UpdatePortafolio(&portafolio); err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to update portafolio",
		})
		return
	}

	utils.Respond(context, utils.APIResponse{
		Status:  http.StatusOK,
		Message: "Portafolio updated successfully",
	})
}

func (controller *PortafolioController) DeletePortafolio(context *gin.Context) {
	id := context.Param("id")

	if err := controller.PortafolioService.DeletePortafolio(uuid.MustParse(id)); err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to delete portafolio",
		})
		return
	}

	utils.Respond(context, utils.APIResponse{
		Status:  http.StatusOK,
		Message: "Portafolio deleted successfully",
	})
}

type CreatePortafolioRequest struct {
	Name         string    `json:"name"`
	Category     string    `json:"category"`
	DataSourceID uuid.UUID `json:"data_source_id"`
	UserID       uuid.UUID `json:"user_id"`
}

func (controller *PortafolioController) CreatePortafolio(context *gin.Context) {
	var createPortafolioRequest CreatePortafolioRequest
	if err := context.ShouldBindJSON(&createPortafolioRequest); err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to create portafolio",
		})
		return
	}

	portafolio := models.Portafolio{
		Name:         createPortafolioRequest.Name,
		Category:     createPortafolioRequest.Category,
		DataSourceID: createPortafolioRequest.DataSourceID,
		UserID:       createPortafolioRequest.UserID,
	}

	if err := controller.PortafolioService.CreatePortafolio(&portafolio); err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to create portafolio",
		})
		return
	}

	utils.Respond(context, utils.APIResponse{
		Status:  http.StatusOK,
		Message: "Portafolio created successfully",
	})
}
