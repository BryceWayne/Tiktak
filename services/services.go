package services

// Services struct encapsulates all the service layer dependencies
// TODO: Implement abstraction layer for services
type Services struct {
	UserService    *UserService
	VideoService   *VideoService
	CommentService *CommentService
	LikeService    *LikeService
}
