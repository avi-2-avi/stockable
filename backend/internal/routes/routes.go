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

	sourceRepo := repositories.NewDataSourceRepository(db)
	sourceService := services.NewDataSourceService(sourceRepo)
	sourceController := controllers.NewDataSourceController(sourceService)

	router.GET("/sources", sourceController.GetSources)

	ratingRepo := repositories.NewAnalystRatingsRepository(db)
	ratingService := services.NewAnalystRatingsService(ratingRepo)
	ratingController := controllers.NewAnalystRatingController(ratingService)

	router.GET("/ratings", ratingController.GetRatings)
	router.GET("/ratings/indicators", ratingController.GetRatingsIndicators)
}

func CORSMiddleware() gin.HandlerFunc {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load configuration")
	}

	allowedOrigin := config.AllowedOrigin
	if allowedOrigin == "" {
		allowedOrigin = "http://localhost:5173"
	}

	return func(context *gin.Context) {
		origin := context.Request.Header.Get("Origin")

		if origin == allowedOrigin {
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
