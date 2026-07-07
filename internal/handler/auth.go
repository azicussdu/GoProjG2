package handler

import (
	"net/http"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/azicussdu/GoProjG2/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(authSrv *service.AuthService) *AuthHandler {
	return &AuthHandler{service: authSrv}
}

func (ah *AuthHandler) Register(c *gin.Context) {
	var input model.RegisterInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid JSON data", nil))
		return
	}

	ctx := c.Request.Context()
	id, err := ah.service.Register(ctx, input)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var input model.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid JSON data", nil))
		return
	}

	ctx := c.Request.Context()
	token, err := ah.service.Login(ctx, input)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Me JWT kimdiki sony kaitarady
func (ah *AuthHandler) Me(c *gin.Context) {
	userID := c.GetInt("userID")

	ctx := c.Request.Context()
	user, err := ah.service.Me(ctx, userID)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
