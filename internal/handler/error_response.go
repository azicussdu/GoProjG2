package handler

import (
	"errors"
	"net/http"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, err error) { // err = AppError ? other Error
	var appErr *apperrors.AppError

	if errors.As(err, &appErr) {
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
}
