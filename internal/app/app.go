package app

import (
	"net/http"

	"github.com/azicussdu/GoProjG2/internal/auth"
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

	jwtManager := auth.NewJWTManager(cfg.JWT.SecretKey, cfg.JWT.AccessTTL, cfg.JWT.RefreshTTL)

	courseRepo := repository.NewPostgresCourseRepo(db)
	lessonRepo := repository.NewPostgresLessonRepo(db)
	userRepo := repository.NewPostgresUserRepo(db)
	enrollmentRepo := repository.NewPostgresEnrollmentRepo(db)

	courseService := service.NewCourseService(courseRepo, lessonRepo, enrollmentRepo, db)
	lessonService := service.NewLessonService(lessonRepo)
	authService := service.NewAuthService(userRepo, jwtManager)
	enrollmentService := service.NewEnrollmentService(enrollmentRepo, courseRepo)

	courseHandler := handler.NewCourseHandler(courseService)
	lessonHandler := handler.NewLessonHandler(lessonService)
	authHandler := handler.NewAuthHandler(authService)
	enrollmentHandler := handler.NewEnrollmentHandler(enrollmentService)

	router := setupRouter(courseHandler, lessonHandler, authHandler, enrollmentHandler, jwtManager)

	srv := &http.Server{Addr: ":" + cfg.Port, Handler: router}
	return srv.ListenAndServe()
}
