package model

import (
	"strings"
	"time"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"gorm.io/gorm"
)

const minTitleLength = 3

var allowedLevels = map[string]bool{
	"beginner":     true,
	"intermediate": true,
	"advanced":     true,
}

type Course struct {
	ID          int     `json:"id" gorm:"primary_key; auto_increment"` // gorm jazbasa da bolady (ID bolsa)
	Title       string  `json:"title" gorm:"type:varchar(255); not null"`
	Description *string `json:"description,omitempty"`
	Price       int     `json:"price" gorm:"default:0; not null"`
	Level       *string `json:"level,omitempty" gorm:"varchar(50)"`
	IsActive    bool    `json:"is_active" gorm:"default:false; not null"`

	TeacherID int  `json:"teacher_id"` // foreign key | users(id) primary key
	Teacher   User `json:"teacher" gorm:"foreignKey:TeacherID;constraint:OnDelete:SET NULL;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

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
