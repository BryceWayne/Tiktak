package services

import (
    "context"
    "log"

    "github.com/BryceWayne/tiktak/models"
    "github.com/BryceWayne/tiktak/pubsubclient"
    "github.com/BryceWayne/tiktak/repositories"
)

type VideoService interface {
    UploadVideo(ctx context.Context, video *models.Video) error
    GetAllVideos(ctx context.Context) ([]models.Video, error)
    // ... other business logic methods
}

type videoService struct {
    videoRepo     repositories.VideoRepository
    pubSubService *pubsubclient.PubSubService
}

// NewVideoService initializes a new VideoService with the given VideoRepository and PubSubService
func NewVideoService(videoRepo repositories.VideoRepository, pubSubService *pubsubclient.PubSubService) VideoService {
    return &videoService{
        videoRepo:     videoRepo,
        pubSubService: pubSubService,
    }
}

func (s *videoService) UploadVideo(ctx context.Context, video *models.Video) error {
    // Business logic for uploading a video
    err := s.videoRepo.SaveVideo(ctx, video)
    if err != nil {
        log.Printf("ERROR: UploadVideo - failed to save video: %v\n", err)
        return err
    }

    // Publish a message after successful video upload
    messageData := []byte("New video uploaded: " + video.ID)
    err = s.pubSubService.PublishMessage(ctx, "new_video_uploaded", messageData)
    if err != nil {
        log.Printf("ERROR: UploadVideo - failed to publish message: %v\n", err)
        return err
    }

    return nil
}

func (s *videoService) GetAllVideos(ctx context.Context) ([]models.Video, error) {
    // Business logic for getting all videos
    return s.videoRepo.GetVideos(ctx)
}

// ... implementation of other methods
