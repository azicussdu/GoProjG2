package service

import (
	"context"
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

func (as *AuthService) Login(ctx context.Context, input model.LoginInput) (string, error) {
	if err := input.Validate(); err != nil {
		return "", err
	}

	user, err := as.repo.GetByEmail(ctx, input.Email)
	if err != nil {
		return "", apperrors.Unauthorized("invalid email or password", nil)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return "", apperrors.Unauthorized("invalid email or password", err)
	}

	accessToken, err := as.jwtManager.NewAccessToken(user)
	if err != nil {
		return "", apperrors.Internal("failed to create access token", err)
	}

	return accessToken, nil
}

func (as *AuthService) Me(ctx context.Context, userID int) (model.User, error) {
	return as.repo.GetByID(ctx, userID)
}
