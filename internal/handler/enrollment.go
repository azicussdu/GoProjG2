package handler

import (
	"net/http"
	"strconv"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/azicussdu/GoProjG2/internal/service"
	"github.com/gin-gonic/gin"
)

type EnrollmentHandler struct {
	service *service.EnrollmentService
}

func NewEnrollmentHandler(enrollmentSrv *service.EnrollmentService) *EnrollmentHandler {
	return &EnrollmentHandler{service: enrollmentSrv}
}

func currentUser(c *gin.Context) (model.User, error) {
	value, exists := c.Get("auth_user")
	if !exists {
		return model.User{}, apperrors.Unauthorized("authentication required", nil)
	}

	user, ok := value.(model.User)
	if !ok {
		return model.User{}, apperrors.Internal("invalid auth context", nil)
	}

	return user, nil
}

// Enroll godoc
// @Summary      Enroll in a course
// @Description  Enrolls the authenticated student in a course. Requires the "student" role.
// @Tags         enrollments
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Course ID"
// @Success      201  {object}  map[string]int
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      403  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /courses/{id}/enroll [post]
func (eh *EnrollmentHandler) Enroll(c *gin.Context) {
	user, err := currentUser(c)
	if err != nil {
		respondWithError(c, err)
		return
	}

	if user.Role != "student" {
		respondWithError(c, apperrors.Forbidden("only students can enroll in courses", nil))
		return
	}

	courseIDStr := c.Param("id")
	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid course ID parameter", nil))
		return
	}

	ctx := c.Request.Context()
	newID, err := eh.service.Enroll(ctx, user.ID, courseID)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": newID})
}

// Leave godoc
// @Summary      Leave a course
// @Description  Removes the authenticated student's enrollment from a course. Requires the "student" role.
// @Tags         enrollments
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Course ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      403  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /courses/{id}/enroll [delete]
func (eh *EnrollmentHandler) Leave(c *gin.Context) {
	user, err := currentUser(c)
	if err != nil {
		respondWithError(c, err)
		return
	}

	if user.Role != "student" {
		respondWithError(c, apperrors.Forbidden("only students can leave courses", nil))
		return
	}

	courseIDStr := c.Param("id")
	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid course ID parameter", nil))
		return
	}

	ctx := c.Request.Context()
	if err = eh.service.Leave(ctx, user.ID, courseID); err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "left the course"})
}

// GetMyCourses godoc
// @Summary      List my courses
// @Description  Returns the courses the authenticated user is enrolled in
// @Tags         enrollments
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   model.Course
// @Failure      401  {object}  map[string]string
// @Router       /enrollments/me [get]
func (eh *EnrollmentHandler) GetMyCourses(c *gin.Context) {
	user, err := currentUser(c)
	if err != nil {
		respondWithError(c, err)
		return
	}

	ctx := c.Request.Context()
	courses, err := eh.service.GetMyCourses(ctx, user.ID)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, courses)
}
