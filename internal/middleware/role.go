package middleware

import (
	"github.com/gin-gonic/gin"
)

func Role(allowedRoles ...string) gin.HandlerFunc {

	rolesMap := make(map[string]bool, len(allowedRoles))
	for _, role := range allowedRoles {
		rolesMap[role] = true
	}

	return func(c *gin.Context) {

	}
}
