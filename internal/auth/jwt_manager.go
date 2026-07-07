package auth

import (
	"errors"
	"fmt"
	"time"

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

func (jm *JWTManager) NewAccessToken(user model.User) (string, error) {
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
		return "", err
	}

	return signedToken, nil
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

func (jm *JWTManager) parseToken(tokenStr, expectedType string) (*Claims, error) {
	if tokenStr == "" {
		return nil, errors.New("token is required")
	}

	parser := jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	parsedToken, err := parser.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jm.secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*Claims)
	if !ok || !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.TokenType != expectedType {
		return nil, fmt.Errorf("unexpected token type: %s", claims.TokenType)
	}

	return claims, nil
}
