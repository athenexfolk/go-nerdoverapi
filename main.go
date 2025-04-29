package main

import (
	"nerdoverapi/internal/category"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/api/v1")
	category.RegisterRoutes(api)

	r.Run("localhost:3000")
}
