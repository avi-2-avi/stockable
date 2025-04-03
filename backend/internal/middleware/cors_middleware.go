package middleware

import (
	"backend/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load configuration")
	}

	allowedOrigins := map[string]bool{
		"http://localhost:5173":        true,
		"http://localhost:3000":        true,
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
