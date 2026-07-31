package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRole_NoAuthUserInContext(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest(http.MethodGet, "/", nil)

	Role("teacher")(c)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
}

func TestRole_UserWithDisallowedRole(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	c.Set("auth_user", model.User{ID: 1, Role: "student"})

	Role("teacher", "admin")(c)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.True(t, c.IsAborted(), "expected the request to be aborted for a disallowed role")
}

func TestRole_UserWithAllowedRole(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	c.Set("auth_user", model.User{ID: 1, Role: "teacher"})

	Role("teacher", "admin")(c)

	assert.False(t, c.IsAborted(), "expected the request not to be aborted for an allowed role")
}
