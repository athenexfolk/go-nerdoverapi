package main

import (
	"github.com/gin-gonic/gin"
	"nerdoverapi/db"
	"nerdoverapi/internal/category"
)

func main() {
	db.ConnectDatabase()
	r := gin.Default()

	api := r.Group("/api/v1")
	category.RegisterRoutes(api)

	r.Run("localhost:3000")
}
