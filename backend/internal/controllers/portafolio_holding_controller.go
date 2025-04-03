package controllers

import (
	"backend/internal/dtos"
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PortafolioHoldingController struct {
	PortafolioHoldingService services.PortafolioHoldingService
}

func NewPortafolioHoldingController(portafolioHoldingService services.PortafolioHoldingService) *PortafolioHoldingController {
	return &PortafolioHoldingController{
		PortafolioHoldingService: portafolioHoldingService,
	}
}

func (controller *PortafolioHoldingController) GetPortafolioHoldings(context *gin.Context) {
	portafolioID := context.Param("portafolio_id")

	partafolioHoldings, err := controller.PortafolioHoldingService.GetAllByPortafolioID(uuid.MustParse(portafolioID))
	if err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get portafolio holdings",
		})
		return
	}

	var portafolioHoldingsDTOs []dtos.PortafolioHoldingDTO
	for _, portafolioHolding := range partafolioHoldings {
		portafolioHoldingsDTOs = append(portafolioHoldingsDTOs, dtos.PortafolioHoldingDTO{
			ID:           portafolioHolding.ID,
			PortafolioID: portafolioHolding.PortafolioID,
			Company: dtos.CompanyDTO{
				ID:   portafolioHolding.Company.ID,
				Name: portafolioHolding.Company.Name,
			},
			Quantity:      portafolioHolding.Quantity,
			PurchasePrice: portafolioHolding.PurchasePrice,
		})
	}

	utils.Respond(context, utils.APIResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Body:    portafolioHoldingsDTOs,
	})
}

type CreatePortafolioHoldingRequest struct {
	PortafolioID  uuid.UUID `json:"portafolio_id"`
	CompanyID     uuid.UUID `json:"company_id"`
	Quantity      float64   `json:"quantity"`
	PurchasePrice float64   `json:"purchase_price"`
	PurchasedAt   time.Time `json:"purchased_at"`
}

func (controller *PortafolioHoldingController) CreatePortafolioHolding(context *gin.Context) {
	var request CreatePortafolioHoldingRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to create portafolio holding",
		})
		return
	}

	portafolioHolding := models.PortafolioHolding{
		PortafolioID:  request.PortafolioID,
		CompanyID:     request.CompanyID,
		Quantity:      request.Quantity,
		PurchasePrice: request.PurchasePrice,
		PurchasedAt:   request.PurchasedAt,
	}

	if err := controller.PortafolioHoldingService.Create(&portafolioHolding); err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to create portafolio holding",
		})
		return
	}

	utils.Respond(context, utils.APIResponse{
		Status:  http.StatusOK,
		Message: "Portafolio holding created successfully",
	})
}

type UpdatePortafolioHoldingRequest struct {
	Quantity      float64   `json:"quantity"`
	PurchasePrice float64   `json:"purchase_price"`
	PurchasedAt   time.Time `json:"purchased_at"`
}

func (controller *PortafolioHoldingController) UpdatePortafolioHolding(context *gin.Context) {
	id := context.Param("id")

	var request UpdatePortafolioHoldingRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update portafolio holding",
		})
		return
	}

	portafolioHolding := models.PortafolioHolding{ID: uuid.MustParse(id)}

	if err := controller.PortafolioHoldingService.Update(&portafolioHolding); err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to update portafolio holding",
		})
		return
	}

	utils.Respond(context, utils.APIResponse{
		Status:  http.StatusOK,
		Message: "Portafolio holding updated successfully",
	})
}

func (controller *PortafolioHoldingController) DeletePortafolioHolding(context *gin.Context) {
	id := context.Param("id")

	if err := controller.PortafolioHoldingService.Delete(uuid.MustParse(id)); err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to delete portafolio holding",
		})
		return
	}

	utils.Respond(context, utils.APIResponse{
		Status:  http.StatusOK,
		Message: "Portafolio holding deleted successfully",
	})
}
