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
	RoleName string `json:"role_name" binding:"required"`
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

	userDto, err := controller.authService.Register(req.FullName, req.Email, req.Password, req.RoleName)
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
		Body:    userDto,
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

	userDto, err := controller.authService.Login(req.Email, req.Password)
	if err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusUnauthorized,
			Message: "Invalid email or password",
		})
		return
	}

	token, err := utils.GenerateJWT(userDto)
	if err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to generate JWT token",
		})
		return
	}

	utils.Respond(context, utils.APIResponse{
		Status:  http.StatusOK,
		Message: "Login successful",
		Body: map[string]interface{}{
			"token": token,
		},
	})
}

func (controller *AuthController) Logout(context *gin.Context) {
	context.SetCookie("auth_token", "", -1, "/", "", false, true)

	if context.Request.Header.Get("auth_token") != "" {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to log out",
		})
		return
	}

	utils.Respond(context, utils.APIResponse{
		Status:  http.StatusOK,
		Message: "Logged out successfully",
	})
}

func (controller *AuthController) Delete(context *gin.Context) {
	id := context.Param("id")

	if err := controller.authService.Delete(uuid.MustParse(id)); err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error deleting user",
		})
		return
	}

	utils.Respond(context, utils.APIResponse{
		Status:  http.StatusOK,
		Message: "User deleted successfully",
	})
}

type UpdateUserRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (controller *AuthController) Update(context *gin.Context) {
	id := context.Param("id")
	var req UpdateUserRequest

	if err := context.ShouldBindJSON(&req); err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Invalid Request",
		})
		return
	}

	user := &models.User{ID: uuid.MustParse(id)}

	if req.FullName != "" {
		user.FullName = req.FullName
	}

	if req.Email != "" {
		user.Email = req.Email
	}

	if req.Password != "" {
		user.Password = req.Password
	}

	err := controller.authService.Update(user)
	if err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error updating user",
			Body: dtos.UpdateUserDTO{
				ID:       user.ID,
				Email:    user.Email,
				FullName: user.FullName,
			},
		})
		return
	}

	utils.Respond(context, utils.APIResponse{
		Status:  http.StatusOK,
		Message: "User updated successfully",
	})
}

func (controller *AuthController) List(context *gin.Context) {
	userDtos, err := controller.authService.List()

	if err != nil {
		utils.Respond(context, utils.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error listing users",
		})
		return
	}

	utils.Respond(context, utils.APIResponse{
		Status:  http.StatusOK,
		Message: "Users listed successfully",
		Body:    userDtos,
	})
}
