package routes

import (
	"backend/config"
	"backend/internal/controllers"
	"backend/internal/repositories"
	"backend/internal/services"
	"fmt"
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
	router.Use(CORSMiddleware())

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

	ratingRepo := repositories.NewAnalystRatingsRepository(db)
	ratingService := services.NewAnalystRatingsService(ratingRepo)
	ratingController := controllers.NewAnalystRatingController(ratingService)

	api.GET("/ratings", ratingController.GetRatings)
	api.GET("/ratings/indicators", ratingController.GetRatingsIndicators)
}

func CORSMiddleware() gin.HandlerFunc {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load configuration")
	}

	allowedOrigins := map[string]bool{
		"http://localhost:5173":        true,
		"http://stockable-frontend":    true,
		"http://stockable-frontend:80": true,
	}

	if config.AllowedOrigin != "" {
		allowedOrigins[config.AllowedOrigin] = true
	}

	return func(context *gin.Context) {
		origin := context.Request.Header.Get("Origin")

		if allowedOrigins[origin] {
			context.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		context.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if context.Request.Method == http.MethodOptions {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}

		context.Next()
	}
}
