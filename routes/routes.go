package routes

import (
	_ "github.com/BryceWayne/tiktak/docs" // Your Swagger generated docs
	"github.com/BryceWayne/tiktak/handlers"
	"github.com/BryceWayne/tiktak/services"
	fiberSwagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userService services.UserService, videoService services.VideoService, commentService services.CommentService, likeService services.LikeService) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Tiktak 👋")
	})
	app.Get("/swagger/*", fiberSwagger.HandlerDefault)

	app.Post("/register", func(c *fiber.Ctx) error {
		return handlers.RegisterUser(c, userService)
	})
	app.Post("/login", func(c *fiber.Ctx) error {
		return handlers.LoginUser(c, userService)
	})

	// Add routes for video operations
	app.Post("/upload", func(c *fiber.Ctx) error {
		return handlers.UploadVideo(c, videoService)
	})
	app.Get("/videos", func(c *fiber.Ctx) error {
		return handlers.GetVideos(c, videoService)
	})

	// Add routes for comment operations
	app.Post("/comment", func(c *fiber.Ctx) error {
		return handlers.PostComment(c, commentService)
	})
	app.Get("/comments/:videoId", func(c *fiber.Ctx) error {
		return handlers.GetComments(c, commentService)
	})

	// Add routes for like operations
	app.Post("/like", func(c *fiber.Ctx) error {
		return handlers.AddLike(c, likeService)
	})
	app.Get("/likes/:videoId", func(c *fiber.Ctx) error {
		return handlers.GetLikes(c, likeService)
	})
}
