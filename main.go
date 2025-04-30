package main

import (
	"nerdoverapi/db"
	"nerdoverapi/storage"

	"nerdoverapi/internal/category"
	"nerdoverapi/internal/feature"
	"nerdoverapi/internal/lesson"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	db.InitFirestore()
	storage.InitStorage()

	r := gin.Default()

	api := r.Group("/api/v1")
	category.RegisterRoutes(api)
	lesson.RegisterRoutes(api)
	feature.RegisterRoutes(api)

	r.Run()
}
