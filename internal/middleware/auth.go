package middleware

import (
	"net/http"
	"strings"

	"github.com/azicussdu/GoProjG2/internal/auth"
	"github.com/gin-gonic/gin"
)

func Auth(jwtManager *auth.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization") // header = "Bearer AS#$#"
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is empty"})
			return
		}

		// parts[0] = "Bearer"
		// parts[1] = "ASD#3443dF"
		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is invalid"})
			return
		}

		user, err := jwtManager.ParseAccessToken(strings.TrimSpace(parts[1]))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set("auth_user", user)
		c.Next()
	}
}
