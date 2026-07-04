package app

import (
	"github.com/azicussdu/GoProjG2/internal/handler"
	"github.com/gin-gonic/gin"
)

func setupRouter(
	courseHandler *handler.CourseHandler,
	lessonHandler *handler.LessonHandler,
) *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		courses := api.Group("/courses")
		{
			courses.GET("", courseHandler.GetAll)
			courses.GET("/:id", courseHandler.GetByID)
			courses.POST("", courseHandler.Create)
			courses.PUT("/:id", courseHandler.Update)
			courses.DELETE("/:id", courseHandler.Delete)
			courses.GET("/:id/lessons", lessonHandler.GetByCourseID)
		}

		lessons := api.Group("/lessons")
		{
			lessons.GET("/:id", lessonHandler.GetByID)
			lessons.POST("", lessonHandler.Create)
			lessons.PUT("/:id", lessonHandler.Update)
			lessons.DELETE("/:id", lessonHandler.Delete)
		}
	}

	return router
}
