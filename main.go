package main

import (
	"log"
	"nerdoverapi/db"
	"nerdoverapi/storage"
	"os"
	"time"

	"nerdoverapi/internal/auth"
	"nerdoverapi/internal/category"
	"nerdoverapi/internal/feature"
	"nerdoverapi/internal/lesson"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env file not found, continuing without it")
	}

	gin.SetMode(gin.ReleaseMode)

	db.InitFirestore()
	storage.InitStorage()

	r := gin.Default()
	r.RedirectTrailingSlash = false

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, // Angular dev server
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	
	api := r.Group("/api/v1")
	
	api.Use(auth.JWTAuthMiddleware())
	
	category.RegisterRoutes(api)
	lesson.RegisterRoutes(api)
	feature.RegisterRoutes(api)
	auth.RegisterRoutes(api)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
