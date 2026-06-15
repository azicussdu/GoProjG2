package main

import (
	"fmt"
	"net/http"

	"github.com/azicussdu/GoProjG2/internal/config"
	"github.com/azicussdu/GoProjG2/internal/handler"
	"github.com/azicussdu/GoProjG2/internal/repository"
	"github.com/azicussdu/GoProjG2/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		return
	}

	db, err := repository.NewPostgresDB(cfg)

	courseRepo := repository.NewPostgresCourseRepo(db)
	courseService := service.NewCourseService(courseRepo)
	courseHandler := handler.NewCourseHandler(courseService)

	router := gin.New()

	router.GET("/api/courses", courseHandler.GetAll)
	router.GET("/api/courses/:id", courseHandler.GetByID)
	router.DELETE("/api/courses/:id", courseHandler.Delete)
	router.POST("/api/courses", courseHandler.Create)
	router.PUT("/api/courses/:id", courseHandler.Update)

	srv := &http.Server{Addr: ":" + cfg.Port, Handler: router}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("Server error")
	}
}
