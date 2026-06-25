package model

import (
	"strings"
	"time"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
)

const minTitleLength = 3

var allowedLevels = map[string]bool{
	"beginner":     true,
	"intermediate": true,
	"advanced":     true,
}

type Course struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description *string `json:"description,omitempty" db:"description"`
	Price       int     `json:"price" db:"price"`
	Level       *string `json:"level,omitempty" db:"level"`
	IsActive    bool    `json:"is_active" db:"is_active"`
	TeacherID   int     `json:"teacher_id" db:"teacher_id"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Create
type CreateCourse struct {
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description"`
	Price       int     `json:"price"`
	Level       *string `json:"level"`
	IsActive    bool    `json:"is_active"`
	TeacherID   int     `json:"teacher_id" binding:"required"`
}

func (cc *CreateCourse) Validate() error {
	cc.Title = strings.TrimSpace(cc.Title)
	if len(cc.Title) < minTitleLength {
		return apperrors.BadRequest("course title is too short", nil)
	}

	if cc.Price < 0 {
		return apperrors.BadRequest("price can not be negative", nil)
	}

	if cc.TeacherID <= 0 {
		return apperrors.BadRequest("invalid teacher ID", nil)
	}

	if cc.Level != nil && !allowedLevels[*cc.Level] {
		return apperrors.BadRequest("invalid course level", nil)
	}

	return nil
}

// Update
type UpdateCourse struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Price       *int    `json:"price"`
	Level       *string `json:"level"`
	IsActive    *bool   `json:"is_active"`
	TeacherID   *int    `json:"teacher_id"`
}

func (uc *UpdateCourse) Validate() error {
	if uc.Title != nil {
		*uc.Title = strings.TrimSpace(*uc.Title)
		if len(*uc.Title) < minTitleLength {
			return apperrors.BadRequest("course title is too short", nil)
		}
	}

	if uc.Price != nil && *uc.Price < 0 {
		return apperrors.BadRequest("price can not be negative", nil)
	}

	if uc.TeacherID != nil && *uc.TeacherID <= 0 {
		return apperrors.BadRequest("invalid teacher ID", nil)
	}

	if uc.Level != nil && !allowedLevels[*uc.Level] {
		return apperrors.BadRequest("invalid course level", nil)
	}

	return nil
}
