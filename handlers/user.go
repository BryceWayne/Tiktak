package handlers

import (
	"github.com/BryceWayne/tiktak/models"
	"github.com/BryceWayne/tiktak/services"
	"github.com/gofiber/fiber/v2"
)

// RegisterUser registers a new user.
// @Summary Register a new user
// @Description Register a user with username, email, and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User required "User Registration Data"
// @Router /register [post]
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

// LoginUser authenticates a user.
// @Summary Authenticate a user
// @Description Authenticate a user and return a token
// @Tags users
// @Accept json
// @Produce json
// @Router /login [post]
func LoginUser(c *fiber.Ctx, userService services.UserService) error {
	// TODO: Implement
	return nil
}
