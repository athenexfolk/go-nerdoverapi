package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup) {
	group := r.Group("/auth")

	group.POST("/", LoginWithGoogleHandler)
}