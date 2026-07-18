package model

import "time"

type Enrollment struct {
	ID        int       `json:"id" db:"id"`
	StudentID int       `json:"student_id" db:"student_id"`
	CourseID  int       `json:"course_id" db:"course_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
