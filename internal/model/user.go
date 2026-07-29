package model

import (
	"strings"
	"time"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
)

const minPasswordLength = 6

var allowedRoles = map[string]bool{
	"student": true,
	"teacher": true,
	"admin":   true,
}

type User struct {
	ID           int       `json:"id" gorm:"primary_key; auto_increment"`
	FullName     string    `json:"full_name" gorm:"type:varchar(255); not null"`
	Email        string    `json:"email" gorm:"type:varchar(255); unique; not null"`
	PasswordHash string    `json:"-" gorm:"not null"`
	Role         string    `json:"role" gorm:"type:varchar(50); default:student; not null"`
	IsActive     bool      `json:"is_active" gorm:"default:true; not null"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type RegisterInput struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}

func (ri *RegisterInput) Validate() error {
	ri.FullName = strings.TrimSpace(ri.FullName)
	ri.Email = strings.TrimSpace(strings.ToLower(ri.Email))

	if ri.FullName == "" {
		return apperrors.BadRequest("name must not be empty", nil)
	}

	if !strings.Contains(ri.Email, "@") {
		return apperrors.BadRequest("invalid email", nil)
	}

	if len(ri.Password) < minPasswordLength {
		return apperrors.BadRequest("password must be at least 6 characters", nil)
	}

	if ri.Role == "" {
		ri.Role = "student"
	}
	if !allowedRoles[ri.Role] {
		return apperrors.BadRequest("invalid role", nil)
	}

	return nil
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (li *LoginInput) Validate() error {
	li.Email = strings.TrimSpace(strings.ToLower(li.Email))

	if li.Email == "" || li.Password == "" {
		return apperrors.BadRequest("email and password are required", nil)
	}
	return nil
}

type RefreshInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type AuthTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessExpAt  int64  `json:"access_at"`
	RefreshExpAt int64  `json:"refresh_at"`
}
