package services

import (
    "context"

    "github.com/BryceWayne/tiktak/models"
    "github.com/BryceWayne/tiktak/repositories"
)

type CommentService interface {
    PostComment(ctx context.Context, comment *models.Comment) error
    GetComments(ctx context.Context, videoID string) ([]models.Comment, error)
    // ... other business logic methods
}

type commentService struct {
    commentRepo repositories.CommentRepository
}

func NewCommentService(commentRepo repositories.CommentRepository) CommentService {
    return &commentService{commentRepo}
}

func (s *commentService) PostComment(ctx context.Context, comment *models.Comment) error {
    // Business logic for posting a comment
    return s.commentRepo.SaveComment(ctx, comment)
}

func (s *commentService) GetComments(ctx context.Context, videoID string) ([]models.Comment, error) {
    // Business logic for getting comments for a video
    return s.commentRepo.GetCommentsByVideoID(ctx, videoID)
}

// ... implementation of other methods
