package model

import (
	"strings"
	"time"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"gorm.io/gorm"
)

type Lesson struct {
	ID        int            `json:"id" gorm:"primary_key; auto_increment"`
	CourseID  int            `json:"course_id" gorm:"not null"`
	Title     string         `json:"title" gorm:"type:varchar(255); not null"`
	Content   *string        `json:"content,omitempty"`
	Position  int            `json:"position" gorm:"default:1; not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type CreateLesson struct {
	CourseID int     `json:"course_id" binding:"required"`
	Title    string  `json:"title" binding:"required"`
	Content  *string `json:"content"`
	Position int     `json:"position"`
}

func (cl *CreateLesson) Validate() error {
	cl.Title = strings.TrimSpace(cl.Title)
	if len(cl.Title) < minTitleLength {
		return apperrors.BadRequest("lesson title is too short", nil)
	}

	if cl.CourseID <= 0 {
		return apperrors.BadRequest("invalid course ID", nil)
	}

	if cl.Position < 0 {
		return apperrors.BadRequest("position cannot be negative", nil)
	}

	return nil
}

type UpdateLesson struct {
	Title    *string `json:"title"`
	Content  *string `json:"content"`
	Position *int    `json:"position"`
}

func (ul *UpdateLesson) Validate() error {
	if ul.Title != nil {
		*ul.Title = strings.TrimSpace(*ul.Title)
		if len(*ul.Title) < minTitleLength {
			return apperrors.BadRequest("lesson title is too short", nil)
		}
	}

	if ul.Position != nil && *ul.Position < 0 {
		return apperrors.BadRequest("position cannot be negative", nil)
	}

	return nil
}
