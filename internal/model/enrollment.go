package model

import (
	"time"

	"gorm.io/gorm"
)

type Enrollment struct {
	ID        int            `json:"id" gorm:"primary_key; auto_increment"`
	StudentID int            `json:"student_id" gorm:"not null"`
	CourseID  int            `json:"course_id" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
