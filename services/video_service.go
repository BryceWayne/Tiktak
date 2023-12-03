package services

import (
    "context"

    "github.com/BryceWayne/tiktak/models"
    "github.com/BryceWayne/tiktak/repositories"
)

type VideoService interface {
    UploadVideo(ctx context.Context, video *models.Video) error
    GetAllVideos(ctx context.Context) ([]models.Video, error)
    // ... other business logic methods
}

type videoService struct {
    videoRepo repositories.VideoRepository
}

func NewVideoService(videoRepo repositories.VideoRepository) VideoService {
    return &videoService{videoRepo}
}

func (s *videoService) UploadVideo(ctx context.Context, video *models.Video) error {
    // Business logic for uploading a video
    return s.videoRepo.SaveVideo(ctx, video)
}

func (s *videoService) GetAllVideos(ctx context.Context) ([]models.Video, error) {
    // Business logic for getting all videos
    return s.videoRepo.GetVideos(ctx)
}

// ... implementation of other methods
