package model

import "time"

type Course struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Price    int    `json:"price"`
	IsActive bool   `json:"is_active"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
