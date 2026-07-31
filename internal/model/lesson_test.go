package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateLesson_Validate(t *testing.T) {
	tests := []struct {
		name    string
		input   CreateLesson
		wantErr bool
	}{
		{
			name:    "valid lesson",
			input:   CreateLesson{CourseID: 1, Title: "Intro", Position: 1},
			wantErr: false,
		},
		{
			name:    "title too short",
			input:   CreateLesson{CourseID: 1, Title: "Hi", Position: 1},
			wantErr: true,
		},
		{
			name:    "missing course ID",
			input:   CreateLesson{CourseID: 0, Title: "Intro", Position: 1},
			wantErr: true,
		},
		{
			name:    "negative position",
			input:   CreateLesson{CourseID: 1, Title: "Intro", Position: -1},
			wantErr: true,
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

func TestUpdateLesson_Validate(t *testing.T) {
	tests := []struct {
		name    string
		input   UpdateLesson
		wantErr bool
	}{
		{"no fields set", UpdateLesson{}, false},
		{"valid title update", UpdateLesson{Title: strPtr("New Title")}, false},
		{"title too short", UpdateLesson{Title: strPtr("Hi")}, true},
		{"negative position", UpdateLesson{Position: intPtr(-1)}, true},
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
