package service

import (
	"context"
	"strings"
	"time"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/azicussdu/GoProjG2/internal/auth"
	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/azicussdu/GoProjG2/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo       repository.UserRepoI
	jwtManager *auth.JWTManager
}

func NewAuthService(userRepo repository.UserRepoI, jwtManager *auth.JWTManager) *AuthService {
	return &AuthService{
		repo:       userRepo,
		jwtManager: jwtManager,
	}
}

func (as *AuthService) Register(ctx context.Context, input model.RegisterInput) (int, error) {
	if err := input.Validate(); err != nil {
		return 0, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, apperrors.Internal("failed to hash password", err)
	}

	user := model.User{
		FullName:     input.FullName,
		Email:        input.Email,
		PasswordHash: string(passwordHash),
		Role:         input.Role,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return as.repo.Create(ctx, user)
}

func (as *AuthService) Login(ctx context.Context, input model.LoginInput) (model.AuthTokens, error) {
	if err := input.Validate(); err != nil {
		return model.AuthTokens{}, err
	}

	user, err := as.repo.GetByEmail(ctx, input.Email)
	if err != nil {
		return model.AuthTokens{}, apperrors.Unauthorized("invalid email or password", nil)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return model.AuthTokens{}, apperrors.Unauthorized("invalid email or password", err)
	}

	accessToken, accessExp, err := as.jwtManager.NewAccessToken(user)
	if err != nil {
		return model.AuthTokens{}, apperrors.Internal("failed to create access token", err)
	}

	refreshToken, refreshExp, err := as.jwtManager.NewRefreshToken(user)
	if err != nil {
		return model.AuthTokens{}, apperrors.Internal("failed to create refresh token", err)
	}

	return model.AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		AccessExpAt:  accessExp,
		RefreshExpAt: refreshExp,
	}, nil
}

func (as *AuthService) Refresh(ctx context.Context, refreshToken string) (model.AuthTokens, error) {
	refreshToken = strings.TrimSpace(refreshToken)
	if refreshToken == "" {
		return model.AuthTokens{}, apperrors.BadRequest("empty refresh token", nil)
	}

	user, err := as.jwtManager.ParseRefreshToken(refreshToken)
	if err != nil {
		return model.AuthTokens{}, apperrors.Unauthorized("invalid refresh token", err)
	}

	accessToken, accessExp, err := as.jwtManager.NewAccessToken(*user)
	if err != nil {
		return model.AuthTokens{}, apperrors.Internal("failed to create access token", err)
	}

	refreshToken, refreshExp, err := as.jwtManager.NewRefreshToken(*user)
	if err != nil {
		return model.AuthTokens{}, apperrors.Internal("failed to create refresh token", err)
	}

	return model.AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		AccessExpAt:  accessExp,
		RefreshExpAt: refreshExp,
	}, nil
}

func (as *AuthService) Me(ctx context.Context, userID int) (model.User, error) {
	return as.repo.GetByID(ctx, userID)
}
