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

	RegisterNoAuthRoutes(api, db)
	RegisterAdminRoutes(api, db)
	RegisterUserRoutes(api, db)
}

func RegisterNoAuthRoutes(api *gin.RouterGroup, db *gorm.DB) {
	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authService)

	api.POST("/auth/register", authController.Register)
	api.POST("/auth/login", authController.Login)
	api.POST("/auth/logout", authController.Logout)
}

func RegisterAdminRoutes(api *gin.RouterGroup, db *gorm.DB) {
	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authService)

	registerAdminRoute(api, "GET", "/auth/list", authController.List)
	registerAdminRoute(api, "PATCH", "/auth/update/:id", authController.Update)
	registerAdminRoute(api, "DELETE", "/auth/delete/:id", authController.Delete)
}

func registerAdminRoute(api *gin.RouterGroup, method, route string, controllerFunc gin.HandlerFunc) {
	api.Handle(method, route, middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllerFunc)
}

func RegisterUserRoutes(api *gin.RouterGroup, db *gorm.DB) {
	sourceRepo := repositories.NewDataSourceRepository(db)
	sourceService := services.NewDataSourceService(sourceRepo)
	sourceController := controllers.NewDataSourceController(sourceService)

	registerUserRoute(api, "GET", "/sources", sourceController.GetSources)
	registerUserRoute(api, "PATCH", "/sources/:id", sourceController.UpdateSource)

	ratingRepo := repositories.NewAnalystRatingRepository(db)
	ratingService := services.NewAnalystRatingService(ratingRepo)
	ratingController := controllers.NewAnalystRatingController(ratingService)

	registerUserRoute(api, "GET", "/ratings", ratingController.GetRatings)
	registerUserRoute(api, "GET", "/ratings/indicators", ratingController.GetRatingsIndicators)
	registerUserRoute(api, "GET", "/ratings/dashboard", ratingController.GetDashboardRatings)
}

func registerUserRoute(api *gin.RouterGroup, method, route string, controllerFunc gin.HandlerFunc) {
	api.Handle(method, route, middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2), controllerFunc)
}
