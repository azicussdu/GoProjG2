package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func strPtr(s string) *string { return &s }
func intPtr(i int) *int       { return &i }

func TestCreateCourse_Validate(t *testing.T) {
	tests := []struct {
		name    string
		input   CreateCourse
		wantErr bool
	}{
		{
			name:    "valid course",
			input:   CreateCourse{Title: "Go Basics", Price: 1000, TeacherID: 1},
			wantErr: false,
		},
		{
			name:    "title too short",
			input:   CreateCourse{Title: "Go", Price: 1000, TeacherID: 1},
			wantErr: true,
		},
		{
			name:    "negative price",
			input:   CreateCourse{Title: "Go Basics", Price: -1, TeacherID: 1},
			wantErr: true,
		},
		{
			name:    "missing teacher ID",
			input:   CreateCourse{Title: "Go Basics", Price: 1000, TeacherID: 0},
			wantErr: true,
		},
		{
			name:    "invalid level",
			input:   CreateCourse{Title: "Go Basics", Price: 1000, TeacherID: 1, Level: strPtr("expert")},
			wantErr: true,
		},
		{
			name:    "valid level",
			input:   CreateCourse{Title: "Go Basics", Price: 1000, TeacherID: 1, Level: strPtr("beginner")},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// UpdateCourse uses pointer fields so only fields the caller actually
// set are validated. An empty UpdateCourse{} must always be valid.
func TestUpdateCourse_Validate(t *testing.T) {
	tests := []struct {
		name    string
		input   UpdateCourse
		wantErr bool
	}{
		{"no fields set", UpdateCourse{}, false},
		{"valid title update", UpdateCourse{Title: strPtr("New Title")}, false},
		{"title too short", UpdateCourse{Title: strPtr("Hi")}, true},
		{"negative price", UpdateCourse{Price: intPtr(-5)}, true},
		{"invalid teacher ID", UpdateCourse{TeacherID: intPtr(0)}, true},
		{"invalid level", UpdateCourse{Level: strPtr("expert")}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
