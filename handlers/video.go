package handlers

import (
    "context"
    "os"

    "cloud.google.com/go/pubsub"
    "github.com/BryceWayne/tiktak/models"
    "github.com/BryceWayne/tiktak/services"
    "github.com/gofiber/fiber/v2"
)

// UploadVideo uploads a new video.
// @Summary Upload a new video
// @Description Uploads a video to the platform
// @Tags videos
// @Accept json
// @Produce json
// @Param video body models.Video required "Video Details"
// @Router /videos [post]
func UploadVideo(c *fiber.Ctx, videoService services.VideoService) error {
    var video models.Video

    ctx := context.Background()

    if err := c.BodyParser(&video); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    if err := videoService.UploadVideo(c.Context(), &video); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to upload video"})
    }

    // Initialize a Pub/Sub client
    projectID := os.Getenv("GCP_PROJECT_ID")
    client, err := pubsub.NewClient(ctx, projectID)
    if err != nil {
        return err
    }
    defer client.Close()

    // Reference to the new_video_uploaded topic
    topic := client.Topic("new_video_uploaded")

    // Publish a message
    _, err = topic.Publish(ctx, &pubsub.Message{
        Data: []byte("Video uploaded with ID: " + video.ID),
    }).Get(ctx)
    if err != nil {
        return err
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Video uploaded successfully"})
}

// GetVideos retrieves all videos.
// @Summary Get all videos
// @Description Retrieves all videos from the platform
// @Tags videos
// @Produce json
// @Router /videos [get]
func GetVideos(c *fiber.Ctx, videoService services.VideoService) error {
    videos, err := videoService.GetAllVideos(c.Context())
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve videos"})
    }

    return c.JSON(videos)
}

// ... other video-related handlers ...
