package handlers

import (
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

    if err := c.BodyParser(&video); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    if err := videoService.UploadVideo(c.Context(), &video); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to upload video"})
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
