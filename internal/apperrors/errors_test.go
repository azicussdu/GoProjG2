package apperrors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// This is the simplest possible unit test: call a pure function,
// check what it returns. No mocks, no setup, no dependencies.
func TestNotFound(t *testing.T) {
	err := NotFound("course not found", nil)

	assert.Equal(t, http.StatusNotFound, err.Code)
	assert.Equal(t, "course not found", err.Message)
}

func TestAppError_Error(t *testing.T) {
	err := New(http.StatusTeapot, "I am a teapot", nil)

	assert.Equal(t, "I am a teapot", err.Error())
}

func TestNew_WrapsUnderlyingError(t *testing.T) {
	underlying := errors.New("db connection refused")
	err := Internal("failed to save user", underlying)

	assert.ErrorIs(t, err.Err, underlying)
}

// A table-driven test: one test function, many cases. This is the
// idiomatic Go way to test several inputs without copy-pasting the
// same assertions over and over.
func TestConstructors_ReturnExpectedCode(t *testing.T) {
	tests := []struct {
		name     string
		build    func(msg string, err error) *AppError
		wantCode int
	}{
		{"NotFound", NotFound, http.StatusNotFound},
		{"BadRequest", BadRequest, http.StatusBadRequest},
		{"Internal", Internal, http.StatusInternalServerError},
		{"Unauthorized", Unauthorized, http.StatusUnauthorized},
		{"Forbidden", Forbidden, http.StatusForbidden},
		{"Conflict", Conflict, http.StatusConflict},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.build("some message", nil)
			assert.Equal(t, tt.wantCode, err.Code)
		})
	}
}
