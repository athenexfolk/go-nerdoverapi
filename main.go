package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"nerdoverapi/db"
	categoryRouter "nerdoverapi/internal/category/router"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	db.InitFirestore()

	r := gin.Default()

	api := r.Group("/api/v1")
	categoryRouter.RegisterRoutes(api)

	r.Run("localhost:3000")
}
