package main

import (
	"fmt"
	"log"
	"os"

	"go-api/config"
	"go-api/models"
	"go-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect DB
	config.ConnectDB()

	// Migrate model
	config.DB.AutoMigrate(&models.User{})

	// Init Fiber
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}
}
