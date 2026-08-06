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

// Register godoc
// @Summary      Register a new user
// @Description  Creates a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      model.RegisterInput  true  "Registration data"
// @Success      201    {object}  map[string]int
// @Failure      400    {object}  map[string]string
// @Failure      409    {object}  map[string]string
// @Router       /auth/register [post]
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

// Login godoc
// @Summary      Log in
// @Description  Authenticates a user and returns access/refresh tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      model.LoginInput  true  "Login credentials"
// @Success      200    {object}  model.AuthTokens
// @Failure      400    {object}  map[string]string
// @Failure      401    {object}  map[string]string
// @Router       /auth/login [post]
func (ah *AuthHandler) Login(c *gin.Context) {
	var input model.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid JSON data", nil))
		return
	}

	ctx := c.Request.Context()
	authTokens, err := ah.service.Login(ctx, input)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, authTokens)
}

// Me godoc
// @Summary      Get current user
// @Description  Returns the profile of the currently authenticated user (resolved from the JWT access token)
// @Tags         auth
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  model.User
// @Failure      401  {object}  map[string]string
// @Router       /auth/me [get]
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

// Refresh godoc
// @Summary      Refresh access token
// @Description  Issues a new pair of access/refresh tokens using a valid refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      model.RefreshInput  true  "Refresh token"
// @Success      200    {object}  model.AuthTokens
// @Failure      400    {object}  map[string]string
// @Failure      401    {object}  map[string]string
// @Router       /auth/refresh [post]
func (ah *AuthHandler) Refresh(c *gin.Context) {
	var input model.RefreshInput
	if err := c.ShouldBindJSON(&input); err != nil {
		respondWithError(c, apperrors.BadRequest("empty refresh token", nil))
		return
	}

	newTokens, err := ah.service.Refresh(c.Request.Context(), input.RefreshToken)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, newTokens)
}
