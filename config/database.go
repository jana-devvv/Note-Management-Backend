package config

import (
	"fmt"
	"jasen-dev/jd-note/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDB() {
	// Setup Connection Database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s  port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// Migrate database
	err := DB.AutoMigrate(&models.Note{})
	if err != nil {
		log.Fatal("Failed to migrate database")
	}
}
