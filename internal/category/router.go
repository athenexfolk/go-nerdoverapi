package category

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup) {
	group := r.Group("/categories")

	group.POST("/", CreateCategoryHandler)
	group.GET("/:id", GetCategoryByIDHandler)
	group.PUT("/:id", UpdateCategoryHandler)
	group.DELETE("/:id", DeleteCategoryHandler)
	group.GET("/", GetAllCategoriesHandler)
}
