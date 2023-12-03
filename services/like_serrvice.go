package services

import (
    "context"

    "github.com/BryceWayne/tiktak/models"
    "github.com/BryceWayne/tiktak/repositories"
)

type LikeService interface {
    AddLike(ctx context.Context, like *models.Like) error
    GetLikes(ctx context.Context, videoID string) ([]models.Like, error)
    // ... other business logic methods
}

type likeService struct {
    likeRepo repositories.LikeRepository
}

func NewLikeService(likeRepo repositories.LikeRepository) LikeService {
    return &likeService{likeRepo}
}

func (s *likeService) AddLike(ctx context.Context, like *models.Like) error {
    // Business logic for adding a like
    return s.likeRepo.SaveLike(ctx, like)
}

func (s *likeService) GetLikes(ctx context.Context, videoID string) ([]models.Like, error) {
    // Business logic for getting likes for a video
    return s.likeRepo.GetLikesByVideoID(ctx, videoID)
}

// ... implementation of other methods
