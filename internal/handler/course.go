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

func (ch *CourseHandler) GetAll(c *gin.Context) {
	courses, err := ch.service.GetAll()
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (ch *CourseHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid ID parameter", nil))
		return
	}

	course, err := ch.service.GetByID(id)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, course)
}

func (ch *CourseHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid ID parameter", nil))
		return
	}

	err = ch.service.Delete(id)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course was deleted"})
}

func (ch *CourseHandler) Create(c *gin.Context) {
	var input model.CreateCourse

	err := c.ShouldBindJSON(&input)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid JSON data", nil))
		return
	}

	newId, err := ch.service.Create(input)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": newId})
}

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

	course, err := ch.service.Update(id, input)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, course)
}
