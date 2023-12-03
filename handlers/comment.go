package handlers

import (
    "github.com/BryceWayne/tiktak/models"
    "github.com/BryceWayne/tiktak/services"
    "github.com/gofiber/fiber/v2"
)

// PostComment posts a comment on a video.
// @Summary Post a comment
// @Description Add a comment to a video
// @Tags comments
// @Accept json
// @Produce json
// @Param comment body models.Comment required "Comment Content"
// @Router /comments [post]
func PostComment(c *fiber.Ctx, commentService services.CommentService) error {
    var comment models.Comment

    if err := c.BodyParser(&comment); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    if err := commentService.PostComment(c.Context(), &comment); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to post comment"})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Comment posted successfully"})
}

// GetComments retrieves all comments for a specific video.
// @Summary Get comments
// @Description Retrieve all comments for a specific video
// @Tags comments
// @Produce json
// @Param videoId path string true "Video ID"
// @Router /comments/{videoId} [get]
func GetComments(c *fiber.Ctx, commentService services.CommentService) error {
    videoID := c.Params("videoId")
    comments, err := commentService.GetComments(c.Context(), videoID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve comments"})
    }

    return c.JSON(comments)
}

// ... other comment-related handlers ...
