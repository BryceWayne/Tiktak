package initializers

import (
	"github.com/BryceWayne/tiktak/repositories"
	"github.com/BryceWayne/tiktak/services"
)

func InitializeServices() (services.UserService, services.VideoService, services.CommentService, services.LikeService) {
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)

	videoRepo := repositories.NewVideoRepository()
	videoService := services.NewVideoService(videoRepo)

	commentRepo := repositories.NewCommentRepository()
	commentService := services.NewCommentService(commentRepo)

	likeRepo := repositories.NewLikeRepository()
	likeService := services.NewLikeService(likeRepo)

	return userService, videoService, commentService, likeService
}
