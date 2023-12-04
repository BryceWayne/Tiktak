package main

import (
	"context"
	"log"
	"os"

	"github.com/BryceWayne/tiktak/initializers"
	"github.com/BryceWayne/tiktak/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Initialize services
	ctx := context.Background()
	projectID := os.Getenv("GCP_PROJECT_ID")
	if projectID == "" {
		// log.Fatalf("GCP_PROJECT_ID environment variable not set\n")
		// return
		projectID = "trends-api-404101"
	}

	userService, videoService, commentService, likeService, err := initializers.InitializeServices(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to initialize services: %v\n", err)
		return
	}

	// Setup routes
	routes.SetupRoutes(app, userService, videoService, commentService, likeService)

	// Start the server
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start the server: %v\n", err)
	}
}
