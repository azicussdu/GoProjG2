package middleware

import (
	"net/http"

	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/gin-gonic/gin"
)

func Role(allowedRoles ...string) gin.HandlerFunc {

	rolesMap := make(map[string]bool, len(allowedRoles))
	for _, role := range allowedRoles {
		rolesMap[role] = true
	}

	return func(c *gin.Context) {
		userVal, ok := c.Get("auth_user")
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user is not authorized"})
			return
		}

		user, ok := userVal.(model.User)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid user in context"})
			return
		}

		if !rolesMap[user.Role] {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid role in context"})
			return
		}

		c.Next()
	}
}
