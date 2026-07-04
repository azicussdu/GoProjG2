package model

import (
	"strings"
	"time"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
)

type Lesson struct {
	ID        int       `json:"id" db:"id"`
	CourseID  int       `json:"course_id" db:"course_id"`
	Title     string    `json:"title" db:"title"`
	Content   *string   `json:"content,omitempty" db:"content"`
	Position  int       `json:"position" db:"position"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
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
