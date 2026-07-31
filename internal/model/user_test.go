package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Table-driven test: each case is a row describing an input and
// whether it should pass or fail validation. Add a new row to test
// a new case - no new test function needed.
func TestRegisterInput_Validate(t *testing.T) {
	tests := []struct {
		name    string
		input   RegisterInput
		wantErr bool
	}{
		{
			name: "valid input",
			input: RegisterInput{
				FullName: "Aisha Bekova",
				Email:    "aisha@example.com",
				Password: "secret123",
				Role:     "student",
			},
			wantErr: false,
		},
		{
			name: "empty full name",
			input: RegisterInput{
				FullName: "   ",
				Email:    "aisha@example.com",
				Password: "secret123",
				Role:     "student",
			},
			wantErr: true,
		},
		{
			name: "email missing @",
			input: RegisterInput{
				FullName: "Aisha Bekova",
				Email:    "aisha.example.com",
				Password: "secret123",
				Role:     "student",
			},
			wantErr: true,
		},
		{
			name: "password too short",
			input: RegisterInput{
				FullName: "Aisha Bekova",
				Email:    "aisha@example.com",
				Password: "abc",
				Role:     "student",
			},
			wantErr: true,
		},
		{
			name: "invalid role",
			input: RegisterInput{
				FullName: "Aisha Bekova",
				Email:    "aisha@example.com",
				Password: "secret123",
				Role:     "superadmin",
			},
			wantErr: true,
		},
		{
			name: "empty role defaults to student",
			input: RegisterInput{
				FullName: "Aisha Bekova",
				Email:    "aisha@example.com",
				Password: "secret123",
				Role:     "",
			},
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

func TestRegisterInput_Validate_NormalizesFields(t *testing.T) {
	input := RegisterInput{
		FullName: "  Aisha Bekova  ",
		Email:    "  AISHA@Example.com  ",
		Password: "secret123",
		Role:     "",
	}

	require.NoError(t, input.Validate())

	assert.Equal(t, "Aisha Bekova", input.FullName)
	assert.Equal(t, "aisha@example.com", input.Email)
	assert.Equal(t, "student", input.Role)
}

func TestLoginInput_Validate(t *testing.T) {
	tests := []struct {
		name    string
		input   LoginInput
		wantErr bool
	}{
		{"valid input", LoginInput{Email: "aisha@example.com", Password: "secret123"}, false},
		{"missing email", LoginInput{Email: "", Password: "secret123"}, true},
		{"missing password", LoginInput{Email: "aisha@example.com", Password: ""}, true},
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
