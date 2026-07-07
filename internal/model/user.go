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
	ID           int       `db:"id" json:"id"`
	FullName     string    `db:"full_name" json:"full_name"`
	Email        string    `db:"email" json:"email"`
	PasswordHash string    `db:"password_hash" json:"-"`
	Role         string    `db:"role" json:"role"`
	IsActive     bool      `db:"is_active" json:"is_active"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
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
