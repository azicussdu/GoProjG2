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
