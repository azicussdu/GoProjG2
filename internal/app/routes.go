package app

import (
	_ "github.com/azicussdu/GoProjG2/docs" // !!! KEREK
	"github.com/azicussdu/GoProjG2/internal/auth"
	"github.com/azicussdu/GoProjG2/internal/handler"
	"github.com/azicussdu/GoProjG2/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter(
	courseHandler *handler.CourseHandler,
	lessonHandler *handler.LessonHandler,
	authHandler *handler.AuthHandler,
	enrollmentHandler *handler.EnrollmentHandler,
	jwtManager *auth.JWTManager,
) *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	teacherAdmin := middleware.Role("teacher", "admin")
	adminOnly := middleware.Role("admin")
	authRequired := middleware.Auth(jwtManager)

	api := router.Group("/api")
	{
		courses := api.Group("/courses") //  /api/courses/...
		{
			courses.GET("", courseHandler.GetAll)
			courses.GET("/:id", courseHandler.GetByID)

			courses.POST("", authRequired, teacherAdmin, courseHandler.Create)
			courses.PUT("/:id", authRequired, teacherAdmin, courseHandler.Update)
			courses.DELETE("/:id", authRequired, adminOnly, courseHandler.Delete)

			courses.GET("/:id/lessons", lessonHandler.GetByCourseID)

			courses.POST("/:id/enroll", authRequired, enrollmentHandler.Enroll)
			courses.DELETE("/:id/enroll", authRequired, enrollmentHandler.Leave)
		}
		enrollments := api.Group("/enrollments") //  /api/enrollments/...
		{
			enrollments.GET("/me", authRequired, enrollmentHandler.GetMyCourses)
		}
		lessons := api.Group("/lessons") //  /api/lessons/...
		{
			lessons.GET("/:id", lessonHandler.GetByID)
			lessons.POST("", middleware.Auth(jwtManager), lessonHandler.Create)
			lessons.PUT("/:id", middleware.Auth(jwtManager), lessonHandler.Update)
			lessons.DELETE("/:id", middleware.Auth(jwtManager), lessonHandler.Delete)
		}
		authGroup := api.Group("/auth") //   /api/auth/register
		{
			authGroup.POST("/register", authHandler.Register)
			authGroup.POST("/login", authHandler.Login)
			authGroup.POST("/refresh", authHandler.Refresh)
			authGroup.GET("/me", authHandler.Me)
		}

	}

	return router
}
