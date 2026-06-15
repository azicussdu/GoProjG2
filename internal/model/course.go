package model

import "time"

type Course struct {
	ID       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	Price    int    `json:"price" db:"price"`
	IsActive bool   `json:"is_active" db:"is_active"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Create
type CreateCourse struct {
	Title    string `json:"title" binding:"required"`
	Price    int    `json:"price"`
	IsActive bool   `json:"is_active"`
}

// Update
type UpdateCourse struct {
	Title    *string `json:"title"`
	Price    *int    `json:"price"`
	IsActive *bool   `json:"is_active"`
}
