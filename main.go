package main

import (
	"jasen-dev/jd-note/config"
	"jasen-dev/jd-note/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error importing .env file")
	}

	config.ConnectDB()
	// Setup Gin router
	r := gin.Default()

	routes.SetupRoutes(r)

	// Run Server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start the server")
	}
}
