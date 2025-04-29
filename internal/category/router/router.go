package router

import (
	"github.com/gin-gonic/gin"
	"nerdoverapi/internal/category/handler"
)

func RegisterRoutes(r *gin.RouterGroup) {
	group := r.Group("/categories")

	group.POST("/", handler.CreateCategoryHandler)
	group.GET("/:id", handler.GetCategoryByIDHandler)
	group.PUT("/:id", handler.UpdateCategoryHandler)
	group.DELETE("/:id", handler.DeleteCategoryHandler)
	group.GET("/", handler.GetAllCategoriesHandler)
}
