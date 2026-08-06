package handler

import (
	"net/http"
	"strconv"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/azicussdu/GoProjG2/internal/service"
	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	service *service.CourseService
}

func NewCourseHandler(courseSrv *service.CourseService) *CourseHandler {
	return &CourseHandler{service: courseSrv}
}

// GetAll godoc
// @Summary      List courses
// @Description  Returns all courses
// @Tags         courses
// @Produce      json
// @Success      200  {array}   model.Course
// @Failure      500  {object}  map[string]string
// @Router       /courses [get]
func (ch *CourseHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	courses, err := ch.service.GetAll(ctx)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, courses)
}

// GetByID godoc
// @Summary      Get a course by ID
// @Description  Returns a single course by its ID
// @Tags         courses
// @Produce      json
// @Param        id   path      int  true  "Course ID"
// @Success      200  {object}  model.Course
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /courses/{id} [get]
func (ch *CourseHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid ID parameter", nil))
		return
	}

	ctx := c.Request.Context()
	course, err := ch.service.GetByID(ctx, id)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, course)
}

// Delete godoc
// @Summary      Delete a course
// @Description  Deletes a course by its ID. Requires admin role.
// @Tags         courses
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Course ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      403  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /courses/{id} [delete]
func (ch *CourseHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid ID parameter", nil))
		return
	}

	ctx := c.Request.Context()
	err = ch.service.Delete(ctx, id)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course was deleted"})
}

// Create godoc
// @Summary      Create a course
// @Description  Creates a new course. Requires teacher or admin role.
// @Tags         courses
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        input  body      model.CreateCourse  true  "Course data"
// @Success      201    {object}  map[string]int
// @Failure      400    {object}  map[string]string
// @Failure      401    {object}  map[string]string
// @Failure      403    {object}  map[string]string
// @Router       /courses [post]
func (ch *CourseHandler) Create(c *gin.Context) {
	var input model.CreateCourse

	err := c.ShouldBindJSON(&input)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid JSON data", nil))
		return
	}

	ctx := c.Request.Context()
	newId, err := ch.service.Create(ctx, input)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": newId})
}

// Update godoc
// @Summary      Update a course
// @Description  Updates an existing course by its ID. Requires teacher or admin role.
// @Tags         courses
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id     path      int                  true  "Course ID"
// @Param        input  body      model.UpdateCourse   true  "Course fields to update"
// @Success      200    {object}  map[string]int
// @Failure      400    {object}  map[string]string
// @Failure      401    {object}  map[string]string
// @Failure      403    {object}  map[string]string
// @Failure      404    {object}  map[string]string
// @Router       /courses/{id} [put]
func (ch *CourseHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid ID parameter", nil))
		return
	}

	var input model.UpdateCourse
	err = c.ShouldBindJSON(&input)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid JSON data", nil))
		return
	}

	ctx := c.Request.Context()
	updatedId, err := ch.service.Update(ctx, id, input)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": updatedId})
}
