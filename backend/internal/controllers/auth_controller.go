package controllers

import (
	"backend/internal/dtos"
	"backend/internal/services"
	"backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService}
}

type RegisterRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (controller *AuthController) Register(context *gin.Context) {
	var req RegisterRequest

	if err := context.ShouldBindJSON(&req); err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Invalid Request",
		})
		return
	}

	err := controller.authService.Register(req.FullName, req.Email, req.Password)
	if err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error registering user",
		})
		return
	}

	utils.Respond(context, utils.APIResponse{
		Status:  http.StatusCreated,
		Message: "User registered successfully",
	})
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (controller *AuthController) Login(context *gin.Context) {
	var req LoginRequest

	if err := context.ShouldBindJSON(&req); err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Invalid Request",
		})
		return
	}

	user, err := controller.authService.Login(req.Email, req.Password)
	if err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusUnauthorized,
			Message: "Invalid email or password",
		})
		return
	}

	context.SetCookie("auth_token", user.Email, 86400, "/", "", false, true)

	utils.Respond(context, utils.APIResponse{
		Status:  http.StatusOK,
		Message: "Login successful",
		Body: map[string]interface{}{
			"user": dtos.LoginUserDTO{
				ID:       user.ID,
				Email:    user.Email,
				FullName: user.FullName,
			},
		},
	})
}

func (c *AuthController) Logout(context *gin.Context) {
	context.SetCookie("auth_token", "", -1, "/", "", false, true)

	utils.Respond(context, utils.APIResponse{
		Status:  http.StatusOK,
		Message: "Logged out successfully",
	})
}
