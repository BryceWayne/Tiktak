package handlers

import (
    "github.com/BryceWayne/tiktak/models"
    "github.com/BryceWayne/tiktak/services"
    "github.com/gofiber/fiber/v2"
)

// AddLike adds a like to a video.
// @Summary Add a like to a video
// @Description Like a video by a user
// @Tags likes
// @Accept json
// @Produce json
// @Param like body models.Like required "Like Details"
// @Router /likes [post]
func AddLike(c *fiber.Ctx, likeService services.LikeService) error {
    var like models.Like

    if err := c.BodyParser(&like); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    if err := likeService.AddLike(c.Context(), &like); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add like"})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Like added successfully"})
}

// GetLikes retrieves all likes for a specific video.
// @Summary Get likes for a video
// @Description Retrieve all likes for a specific video
// @Tags likes
// @Produce json
// @Param videoId path string true "Video ID"
// @Router /likes/{videoId} [get]
func GetLikes(c *fiber.Ctx, likeService services.LikeService) error {
    videoID := c.Params("videoId")
    likes, err := likeService.GetLikes(c.Context(), videoID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve likes"})
    }

    return c.JSON(likes)
}

// ... other like-related handlers ...
