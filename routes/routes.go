package routes

import (
	_ "github.com/BryceWayne/tiktak/docs" // Your Swagger generated docs
	"github.com/BryceWayne/tiktak/handlers"
	"github.com/BryceWayne/tiktak/services"
	fiberSwagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, services services.Services) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Tiktak 👋")
	})
	app.Get("/swagger/*", fiberSwagger.HandlerDefault)

	app.Post("/register", func(c *fiber.Ctx) error {
		return handlers.RegisterUser(c, services.UserService)
	})
	app.Post("/login", func(c *fiber.Ctx) error {
		return handlers.LoginUser(c, services.UserService)
	})

	// Add routes for video operations
	app.Post("/upload", func(c *fiber.Ctx) error {
		return handlers.UploadVideo(c, services.VideoService)
	})
	app.Get("/videos", func(c *fiber.Ctx) error {
		return handlers.GetVideos(c, services.VideoService)
	})

	// Add routes for comment operations
	app.Post("/comment", func(c *fiber.Ctx) error {
		return handlers.PostComment(c, services.CommentService)
	})
	app.Get("/comments/:videoId", func(c *fiber.Ctx) error {
		return handlers.GetComments(c, services.CommentService)
	})

	// Add routes for like operations
	app.Post("/like", func(c *fiber.Ctx) error {
		return handlers.AddLike(c, services.LikeService)
	})
	app.Get("/likes/:videoId", func(c *fiber.Ctx) error {
		return handlers.GetLikes(c, services.LikeService)
	})
}
