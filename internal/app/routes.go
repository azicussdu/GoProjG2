package app

import (
	"github.com/azicussdu/GoProjG2/internal/auth"
	"github.com/azicussdu/GoProjG2/internal/handler"
	"github.com/azicussdu/GoProjG2/internal/middleware"
	"github.com/gin-gonic/gin"
)

func setupRouter(
	courseHandler *handler.CourseHandler,
	lessonHandler *handler.LessonHandler,
	authHandler *handler.AuthHandler,
	jwtManager *auth.JWTManager,
) *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		courses := api.Group("/courses") //  /api/courses/...
		{
			courses.GET("", courseHandler.GetAll)
			courses.GET("/:id", courseHandler.GetByID)

			courses.POST("",
				middleware.Auth(jwtManager),
				middleware.Role("teacher", "admin"),
				courseHandler.Create)

			courses.PUT("/:id", middleware.Auth(jwtManager), courseHandler.Update)
			courses.DELETE("/:id", middleware.Auth(jwtManager), courseHandler.Delete)
			courses.GET("/:id/lessons", lessonHandler.GetByCourseID)
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
