package lesson

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	group := r.Group("/lessons")

	group.POST("/", CreateLessonHandler)
	group.GET("/:id", GetLessonByIDHandler)
	group.PUT("/:id", UpdateLessonHandler)
	group.PATCH("/:id", UpdateContentHandler)
	group.DELETE("/:id", DeleteCategoryHandler)
	group.GET("/", GetAllCategoriesHandler)
}
