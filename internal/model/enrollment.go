package model

import (
	"time"

	"gorm.io/gorm"
)

type Enrollment struct {
	ID int `json:"id" gorm:"primary_key; auto_increment"`

	StudentID int  `json:"student_id" gorm:"not null; uniqueIndex:idx_student_course"`
	Student   User `json:"student" gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE;"`

	CourseID int    `json:"course_id" gorm:"not null; uniqueIndex:idx_student_course"`
	Course   Course `json:"course" gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
