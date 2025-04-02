package middleware

import (
	"backend/internal/database"
	"backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken, err := c.Cookie("auth_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - No auth token"})
			c.Abort()
			return
		}

		db, dbErr := database.Connect()
		if dbErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return
		}

		var user models.User
		result := db.Where("email = ?", authToken).First(&user)
		if result.Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Invalid user"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
