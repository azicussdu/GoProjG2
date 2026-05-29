package main

import (
	"fmt"
	"net/http"

	"github.com/azicussdu/GoProjG2/internal/handler"
	"github.com/azicussdu/GoProjG2/internal/repository"
	"github.com/azicussdu/GoProjG2/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {

	courseRepo := repository.NewCourseRepo()
	courseService := service.NewCourseService(courseRepo)
	courseHandler := handler.NewCourseHandler(courseService)

	router := gin.New()

	router.GET("/api/courses", courseHandler.GetAll)
	router.GET("/api/courses/:id", courseHandler.GetByID)

	srv := &http.Server{Addr: ":8080", Handler: router}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("Server error")
	}
}
