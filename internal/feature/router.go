package feature

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	group := r.Group("/features")
	
	group.GET("/", ExportLessonHandler)
	group.GET("/images", GetAllImagesHandler)
	group.POST("/images", UploadImageHandler)
}
