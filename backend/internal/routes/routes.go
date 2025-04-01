package routes

import (
	"backend/internal/controllers"
	"backend/internal/middleware"
	"backend/internal/repositories"
	"backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return router
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	router.Use(middleware.CORSMiddleware())

	api := router.Group("/api")

	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authService)

	api.POST("/auth/register", authController.Register)
	api.POST("/auth/login", authController.Login)
	api.POST("/auth/logout", authController.Logout)

	sourceRepo := repositories.NewDataSourceRepository(db)
	sourceService := services.NewDataSourceService(sourceRepo)
	sourceController := controllers.NewDataSourceController(sourceService)

	api.GET("/sources", sourceController.GetSources)

	ratingRepo := repositories.NewAnalystRatingRepository(db)
	ratingService := services.NewAnalystRatingService(ratingRepo)
	ratingController := controllers.NewAnalystRatingController(ratingService)

	api.GET("/ratings", ratingController.GetRatings)
	api.GET("/ratings/indicators", ratingController.GetRatingsIndicators)
}
