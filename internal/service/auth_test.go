package service

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/azicussdu/GoProjG2/internal/auth"
	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func testJWTManager() *auth.JWTManager {
	return auth.NewJWTManager("test-secret", time.Minute, time.Hour)
}

func TestAuthService_Register(t *testing.T) {
	t.Run("invalid input is rejected before hitting the repo", func(t *testing.T) {
		svc := NewAuthService(new(MockUserRepo), testJWTManager())

		_, err := svc.Register(context.Background(), model.RegisterInput{
			FullName: "", Email: "bad", Password: "short",
		})

		require.Equal(t, http.StatusBadRequest, appErrorCode(err))
	})

	t.Run("hashes the password before saving", func(t *testing.T) {
		repo := new(MockUserRepo)
		repo.On("Create", mock.Anything, mock.MatchedBy(func(u model.User) bool {
			return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte("secret123")) == nil
		})).Return(1, nil)
		svc := NewAuthService(repo, testJWTManager())

		_, err := svc.Register(context.Background(), model.RegisterInput{
			FullName: "Aisha Bekova",
			Email:    "aisha@example.com",
			Password: "secret123",
			Role:     "student",
		})

		require.NoError(t, err)
		repo.AssertExpectations(t)
	})
}

func TestAuthService_Login(t *testing.T) {
	hash, err := bcrypt.GenerateFromPassword([]byte("secret123"), bcrypt.DefaultCost)
	require.NoError(t, err, "failed to prepare test password hash")
	existingUser := model.User{ID: 1, Email: "aisha@example.com", PasswordHash: string(hash), Role: "student"}

	t.Run("unknown email", func(t *testing.T) {
		repo := new(MockUserRepo)
		repo.On("GetByEmail", mock.Anything, "nobody@example.com").Return(model.User{}, apperrors.NotFound("user not found", nil))
		svc := NewAuthService(repo, testJWTManager())

		_, err := svc.Login(context.Background(), model.LoginInput{Email: "nobody@example.com", Password: "secret123"})

		require.Equal(t, http.StatusUnauthorized, appErrorCode(err))
		repo.AssertExpectations(t)
	})

	t.Run("wrong password", func(t *testing.T) {
		repo := new(MockUserRepo)
		repo.On("GetByEmail", mock.Anything, existingUser.Email).Return(existingUser, nil)
		svc := NewAuthService(repo, testJWTManager())

		_, err := svc.Login(context.Background(), model.LoginInput{Email: existingUser.Email, Password: "wrong-password"})

		require.Equal(t, http.StatusUnauthorized, appErrorCode(err))
		repo.AssertExpectations(t)
	})

	t.Run("success returns tokens", func(t *testing.T) {
		repo := new(MockUserRepo)
		repo.On("GetByEmail", mock.Anything, existingUser.Email).Return(existingUser, nil)
		svc := NewAuthService(repo, testJWTManager())

		tokens, err := svc.Login(context.Background(), model.LoginInput{Email: existingUser.Email, Password: "secret123"})

		require.NoError(t, err)
		assert.NotEmpty(t, tokens.AccessToken)
		assert.NotEmpty(t, tokens.RefreshToken)
		repo.AssertExpectations(t)
	})
}

func TestAuthService_Refresh(t *testing.T) {
	t.Run("empty token", func(t *testing.T) {
		svc := NewAuthService(new(MockUserRepo), testJWTManager())

		_, err := svc.Refresh(context.Background(), "  ")

		require.Equal(t, http.StatusBadRequest, appErrorCode(err))
	})

	t.Run("garbage token", func(t *testing.T) {
		svc := NewAuthService(new(MockUserRepo), testJWTManager())

		_, err := svc.Refresh(context.Background(), "not-a-real-token")

		require.Equal(t, http.StatusUnauthorized, appErrorCode(err))
	})

	t.Run("success issues a new token pair", func(t *testing.T) {
		jwtManager := testJWTManager()
		svc := NewAuthService(new(MockUserRepo), jwtManager)

		refreshToken, _, err := jwtManager.NewRefreshToken(model.User{ID: 1, Email: "aisha@example.com", Role: "student"})
		require.NoError(t, err, "failed to prepare refresh token")

		tokens, err := svc.Refresh(context.Background(), refreshToken)

		require.NoError(t, err)
		assert.NotEmpty(t, tokens.AccessToken)
		assert.NotEmpty(t, tokens.RefreshToken)
	})
}

func TestAuthService_Me(t *testing.T) {
	repo := new(MockUserRepo)
	repo.On("GetByID", mock.Anything, 7).Return(model.User{ID: 7, Email: "aisha@example.com"}, nil)
	svc := NewAuthService(repo, testJWTManager())

	user, err := svc.Me(context.Background(), 7)

	require.NoError(t, err)
	assert.Equal(t, 7, user.ID)
	repo.AssertExpectations(t)
}
