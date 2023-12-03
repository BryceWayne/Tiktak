package handlers

import (
	"github.com/BryceWayne/tiktak/models"
	"github.com/BryceWayne/tiktak/services"
	"github.com/gofiber/fiber/v2"
)

// swagger:route POST /register users registerUser
// RegisterUser registers a new user.
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//
//	200: userResponse
//	400: errorResponse
//	500: errorResponse
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

// swagger:operation POST /login users loginUser
// ---
// summary: Authenticates a user.
// description: This will authenticate a user and return a token.
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
//   - name: user
//     in: body
//     description: User login credentials
//     required: true
//     schema:
//     "$ref": "#/definitions/UserCredentials"
//
// responses:
//
//	'200':
//	  description: Successful login
//	  schema:
//	    "$ref": "#/definitions/UserToken"
//	'401':
//	  description: Unauthorized - Invalid credentials
//	  schema:
//	    "$ref": "#/definitions/errorResponse"
//	'500':
//	  description: Internal Server Error
//	  schema:
//	    "$ref": "#/definitions/errorResponse"
func LoginUser(c *fiber.Ctx, userService services.UserService) error {
	// TODO: Implement
	return nil
}

// Define the response and other models here, or in a separate file
// swagger:model userResponse
type userResponse struct {
	// The user ID
	// in: body
	ID string `json:"id"`
	// The user's username
	// in: body
	Username string `json:"username"`
	// The user's email
	// in: body
	Email string `json:"email"`
}

// swagger:model errorResponse
type errorResponse struct {
	// The error message
	// in: body
	Message string `json:"message"`
}

// swagger:model UserCredentials
type UserCredentials struct {
	// The user's username
	// in: body
	Username string `json:"username"`
	// The user's password
	// in: body
	Password string `json:"password"`
}

// swagger:model UserToken
type UserToken struct {
	// The JWT token for authenticated user
	// in: body
	Token string `json:"token"`
}
