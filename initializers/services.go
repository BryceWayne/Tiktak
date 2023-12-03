package initializers

import (
	"github.com/BryceWayne/tiktak/repositories"
	"github.com/BryceWayne/tiktak/services"
)

// Initialize all services and return a Services struct
func InitializeServices() Services {
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)

	videoRepo := repositories.NewVideoRepository()
	videoService := services.NewVideoService(videoRepo)

	commentRepo := repositories.NewCommentRepository()
	commentService := services.NewCommentService(commentRepo)

	likeRepo := repositories.NewLikeRepository()
	likeService := services.NewLikeService(likeRepo)

	return Services{
		UserService:    userService,
		VideoService:   videoService,
		CommentService: commentService,
		LikeService:    likeService,
	}
}
