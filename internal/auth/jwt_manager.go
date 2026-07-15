package auth

import (
	"time"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/golang-jwt/jwt/v5" // go get github.com/golang-jwt/jwt/v5
)

type Claims struct {
	UserID               int    `json:"uid"`
	Email                string `json:"email"`
	Role                 string `json:"role"`
	TokenType            string `json:"type"` // access, refresh
	jwt.RegisteredClaims        // exp, iat
}

type JWTManager struct {
	secret     []byte
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewJWTManager(secret string, accessTTL, refreshTTL time.Duration) *JWTManager {
	return &JWTManager{
		secret:     []byte(secret),
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
	}
}

func (jm *JWTManager) NewAccessToken(user model.User) (string, int64, error) {
	expiresAt := time.Now().Add(jm.accessTTL)

	claims := Claims{
		UserID:    user.ID,
		Email:     user.Email,
		Role:      user.Role,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()), // number in seconds
			ExpiresAt: jwt.NewNumericDate(expiresAt),  // number in seconds
		},
	}

	// tokennin 2 boligin jasap beredi (bez podpisi)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jm.secret)
	if err != nil {
		return "", 0, err
	}

	return signedToken, expiresAt.Unix(), nil
}

func (jm *JWTManager) NewRefreshToken(user model.User) (string, int64, error) {
	expiresAt := time.Now().Add(jm.refreshTTL)

	claims := Claims{
		UserID:    user.ID,
		Email:     user.Email,
		Role:      user.Role,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()), // number in seconds
			ExpiresAt: jwt.NewNumericDate(expiresAt),  // number in seconds
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jm.secret)
	if err != nil {
		return "", 0, err
	}

	return signedToken, expiresAt.Unix(), nil
}

func (jm *JWTManager) ParseAccessToken(tokenStr string) (*model.User, error) {
	claims, err := jm.parseToken(tokenStr, "access")
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:    claims.UserID,
		Email: claims.Email,
		Role:  claims.Role,
	}, nil
}

func (jm *JWTManager) ParseRefreshToken(tokenStr string) (*model.User, error) {
	claims, err := jm.parseToken(tokenStr, "refresh")
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:    claims.UserID,
		Email: claims.Email,
		Role:  claims.Role,
	}, nil
}

func (jm *JWTManager) parseToken(tokenStr, expectedType string) (*Claims, error) {
	if tokenStr == "" {
		return nil, apperrors.BadRequest("Token is required", nil)
	}

	claims := &Claims{}

	parser := jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	// barin osy tekseredi (exp, validnost)
	parsedToken, err := parser.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jm.secret, nil
	})

	if err != nil || !parsedToken.Valid {
		return nil, apperrors.Unauthorized("Invalid or expired token", nil)
	}

	if claims.TokenType != expectedType {
		return nil, apperrors.BadRequest("Unexpected token type", nil)
	}

	return claims, nil
}
