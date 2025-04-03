package middleware

import (
	"backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(allowedRoleIDs ...uint) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		u, ok := user.(models.User)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Invalid user data"})
			c.Abort()
			return
		}

		for _, roleID := range allowedRoleIDs {
			if u.RoleID == roleID {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden - Insufficient permissions"})
		c.Abort()
	}
}
