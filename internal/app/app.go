package app

import (
	"net/http"

	"github.com/azicussdu/GoProjG2/internal/config"
	"github.com/azicussdu/GoProjG2/internal/handler"
	"github.com/azicussdu/GoProjG2/internal/repository"
	"github.com/azicussdu/GoProjG2/internal/service"
)

func Run(cfg *config.Config) error {
	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		return err
	}
	defer db.Close()

	courseRepo := repository.NewPostgresCourseRepo(db)
	courseService := service.NewCourseService(courseRepo)
	courseHandler := handler.NewCourseHandler(courseService)

	lessonRepo := repository.NewPostgresLessonRepo(db)
	lessonService := service.NewLessonService(lessonRepo)
	lessonHandler := handler.NewLessonHandler(lessonService)

	router := setupRouter(courseHandler, lessonHandler)

	srv := &http.Server{Addr: ":" + cfg.Port, Handler: router}
	return srv.ListenAndServe()
}
