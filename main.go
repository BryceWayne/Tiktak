package main

import (
	"log"

	"github.com/BryceWayne/tiktak/initializers"
	"github.com/BryceWayne/tiktak/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	userService, videoService, commentService, likeService := initializers.InitializeServices()

	routes.SetupRoutes(app, userService, videoService, commentService, likeService)

	// Start the server
	log.Fatalf("%v\n", app.Listen(":8080"))
}
