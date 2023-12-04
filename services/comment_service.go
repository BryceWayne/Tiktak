package services

import (
    "context"
    "encoding/json"
    "log"

    "github.com/BryceWayne/tiktak/models"
    "github.com/BryceWayne/tiktak/pubsubclient"
    "github.com/BryceWayne/tiktak/repositories"
)

type CommentService interface {
    PostComment(ctx context.Context, comment *models.Comment) error
    GetComments(ctx context.Context, videoID string) ([]models.Comment, error)
    // ... other business logic methods
}

type commentService struct {
    commentRepo   repositories.CommentRepository
    pubSubService *pubsubclient.PubSubService
}

func NewCommentService(commentRepo repositories.CommentRepository, pubSubService *pubsubclient.PubSubService) CommentService {
    return &commentService{
        commentRepo:   commentRepo,
        pubSubService: pubSubService,
    }
}

func (s *commentService) PostComment(ctx context.Context, comment *models.Comment) error {
    // Save the comment using the repository
    err := s.commentRepo.SaveComment(ctx, comment)
    if err != nil {
        log.Printf("ERROR: PostComment - failed to save comment: %v\n", err)
        return err
    }

    // Publish a message to a topic when a new comment is posted
    commentData, err := json.Marshal(comment)
    if err != nil {
        // Optionally handle error (e.g., log it), but decide if you want to return it
        log.Printf("ERROR: PostComment - failed to marshal comment: %v\n", err)
        return err
    }

    // Define your Pub/Sub topic name (e.g., "new-comment-posted")
    topicName := "new-comment-posted"

    // Publish the message
    err = s.pubSubService.PublishMessage(ctx, topicName, commentData)
    if err != nil {
        // Handle the error according to your application's needs
        log.Printf("ERROR: PostComment - failed to publish message: %v\n", err)
        return err
    }

    return nil
}

func (s *commentService) GetComments(ctx context.Context, videoID string) ([]models.Comment, error) {
    // Business logic for getting comments for a video
    return s.commentRepo.GetCommentsByVideoID(ctx, videoID)
}

// ... implementation of other methods
