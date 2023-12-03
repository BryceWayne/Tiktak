package handlers

import (
	"github.com/BryceWayne/tiktak/models"
	"github.com/BryceWayne/tiktak/services"
	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx, userService services.UserService) error {
	// Create a new User struct to hold the user's input
	var user models.User

	// Parse the request body into the user struct
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Call the UserService to register the new user
	if err := userService.RegisterUser(c.Context(), &user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}

	// Return a success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User registered successfully"})
}

func LoginUser(c *fiber.Ctx, userService services.UserService) error {
	// TODO: Implement
	return nil
}
