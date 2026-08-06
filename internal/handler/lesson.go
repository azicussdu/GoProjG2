package handler

import (
	"net/http"
	"strconv"

	"github.com/azicussdu/GoProjG2/internal/apperrors"
	"github.com/azicussdu/GoProjG2/internal/model"
	"github.com/azicussdu/GoProjG2/internal/service"
	"github.com/gin-gonic/gin"
)

type LessonHandler struct {
	service *service.LessonService
}

func NewLessonHandler(lessonSrv *service.LessonService) *LessonHandler {
	return &LessonHandler{service: lessonSrv}
}

// GetByCourseID godoc
// @Summary      List lessons of a course
// @Description  Returns all lessons that belong to the given course
// @Tags         lessons
// @Produce      json
// @Param        id   path      int  true  "Course ID"
// @Success      200  {array}   model.Lesson
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /courses/{id}/lessons [get]
func (h *LessonHandler) GetByCourseID(c *gin.Context) {
	courseIDStr := c.Param("id")
	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid course ID parameter", nil))
		return
	}

	ctx := c.Request.Context()
	lessons, err := h.service.GetByCourseID(ctx, courseID)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, lessons)
}

// GetByID godoc
// @Summary      Get a lesson by ID
// @Description  Returns a single lesson by its ID
// @Tags         lessons
// @Produce      json
// @Param        id   path      int  true  "Lesson ID"
// @Success      200  {object}  model.Lesson
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /lessons/{id} [get]
func (h *LessonHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid ID parameter", nil))
		return
	}

	ctx := c.Request.Context()
	lesson, err := h.service.GetByID(ctx, id)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, lesson)
}

// Create godoc
// @Summary      Create a lesson
// @Description  Creates a new lesson. Requires authentication.
// @Tags         lessons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        input  body      model.CreateLesson  true  "Lesson data"
// @Success      201    {object}  map[string]int
// @Failure      400    {object}  map[string]string
// @Failure      401    {object}  map[string]string
// @Router       /lessons [post]
func (h *LessonHandler) Create(c *gin.Context) {
	var input model.CreateLesson

	if err := c.ShouldBindJSON(&input); err != nil {
		respondWithError(c, apperrors.BadRequest("invalid JSON data", nil))
		return
	}

	ctx := c.Request.Context()
	newID, err := h.service.Create(ctx, input)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": newID})
}

// Update godoc
// @Summary      Update a lesson
// @Description  Updates an existing lesson by its ID. Requires authentication.
// @Tags         lessons
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id     path      int                  true  "Lesson ID"
// @Param        input  body      model.UpdateLesson   true  "Lesson fields to update"
// @Success      200    {object}  map[string]int
// @Failure      400    {object}  map[string]string
// @Failure      401    {object}  map[string]string
// @Failure      404    {object}  map[string]string
// @Router       /lessons/{id} [put]
func (h *LessonHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid ID parameter", nil))
		return
	}

	var input model.UpdateLesson
	if err = c.ShouldBindJSON(&input); err != nil {
		respondWithError(c, apperrors.BadRequest("invalid JSON data", nil))
		return
	}

	ctx := c.Request.Context()
	updatedID, err := h.service.Update(ctx, id, input)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": updatedID})
}

// Delete godoc
// @Summary      Delete a lesson
// @Description  Deletes a lesson by its ID. Requires authentication.
// @Tags         lessons
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Lesson ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /lessons/{id} [delete]
func (h *LessonHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid ID parameter", nil))
		return
	}

	ctx := c.Request.Context()
	if err = h.service.Delete(ctx, id); err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lesson was deleted"})
}
