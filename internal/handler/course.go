package handler

import (
	"net/http"
	"strconv"

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (ch *CourseHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}

	course, err := ch.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course is not found"})
		return
	}

	c.JSON(http.StatusOK, course)
}

func (ch *CourseHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}

	err = ch.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course was deleted"})
}

func (ch *CourseHandler) Create(c *gin.Context) {
	var input model.CreateCourse

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad json data"})
		return
	}

	newId, err := ch.service.Create(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Some error happened"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": newId})
}

func (ch *CourseHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}

	var input model.UpdateCourse
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad json data"})
		return
	}

	course, err := ch.service.Update(id, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong fields"})
		return
	}

	c.JSON(http.StatusOK, course)
}
