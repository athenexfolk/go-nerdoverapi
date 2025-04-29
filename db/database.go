package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"nerdoverapi/internal/category"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("nerdover.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	// Auto-migrate the Category model
	err = database.AutoMigrate(&category.Category{})
	if err != nil {
		log.Fatal("Failed to migrate database!", err)
	}

	DB = database
}
