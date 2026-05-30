package model

import "time"

// GetAll, GetByID
type Course struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Price    int    `json:"price"`
	IsActive bool   `json:"is_active"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
