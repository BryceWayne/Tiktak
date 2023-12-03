package handlers

import (
    "github.com/BryceWayne/tiktak/models"
    "github.com/BryceWayne/tiktak/services"
    "github.com/gofiber/fiber/v2"
)

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

func GetComments(c *fiber.Ctx, commentService services.CommentService) error {
    videoID := c.Params("videoId")
    comments, err := commentService.GetComments(c.Context(), videoID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve comments"})
    }

    return c.JSON(comments)
}

// ... other comment-related handlers ...
