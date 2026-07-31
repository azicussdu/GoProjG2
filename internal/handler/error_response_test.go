package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// newTestContext builds a minimal *gin.Context backed by an
// httptest.ResponseRecorder, so we can inspect exactly what status
// code and body a handler function writes - without starting a real
// HTTP server.
func newTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	return c, recorder
}

func TestRespondWithError_AppError(t *testing.T) {
	c, recorder := newTestContext()

	respondWithError(c, apperrors.NotFound("course not found", nil))

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "course not found")
}

func TestRespondWithError_UnknownError(t *testing.T) {
	c, recorder := newTestContext()

	respondWithError(c, errors.New("some unexpected failure"))

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	// The raw error message must never leak to the client - only the
	// generic "internal server error" message should appear.
	assert.NotContains(t, recorder.Body.String(), "some unexpected failure")
}
