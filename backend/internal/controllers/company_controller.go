package controllers

import (
	"backend/internal/services"
	"backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyController struct {
	CompanyService services.CompanyService
}

func NewCompanyController(companyService services.CompanyService) *CompanyController {
	return &CompanyController{
		CompanyService: companyService,
	}
}

func (controller *CompanyController) GetCompanyDescription(context *gin.Context) {
	ticker := context.Query("ticker")
	company := context.Query("company")

	description, err := controller.CompanyService.GetCompanyDescription(ticker, company)
	if err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get company description",
		})
		return
	}

	utils.Respond(context, utils.APIResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Body: gin.H{
			"description": description,
		},
	})
}
