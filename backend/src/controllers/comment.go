package controllers

import (
	"shared-blog-backend/src/database"
	"shared-blog-backend/src/models"
	"shared-blog-backend/src/repositories"
	"shared-blog-backend/src/responses"
	"shared-blog-backend/src/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CommentPost(c *fiber.Ctx) error {

	var comment models.Comment

	if err := c.BodyParser(&comment); err != nil {
		return responses.Error(c, fiber.StatusUnprocessableEntity, fiber.Map{
			"message": "error while reading body",
		})
	}

	if err := comment.ValidateCommentFields(); err != nil {
		return responses.Error(c, fiber.StatusUnprocessableEntity, fiber.Map{
			"message": err.Error(),
		})
	}

	db, err := database.Connect()

	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	defer db.Close()

	userOnToken, _ := utils.IsTokenValid(c)

	repo := repositories.CommentRepository(db)

	if err = repo.CommentPost(userOnToken, comment); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	return responses.SendJSON(c, fiber.StatusCreated, fiber.Map{
		"message": "success",
	})
}

func DeleteComment(c *fiber.Ctx) error {

	params := c.Params("commentId")

	commentId, err := strconv.ParseUint(params, 10, 64)

	if err != nil {
		return responses.Error(c, fiber.StatusBadRequest, fiber.Map{
			"message": err.Error(),
		})
	}

	db, err := database.Connect()

	if err != nil {
		return responses.SendJSON(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	defer db.Close()

	userOnToken, _ := utils.IsTokenValid(c)

	repo := repositories.CommentRepository(db)
	status, err := repo.DeleteComment(userOnToken, commentId)

	if err != nil {
		return responses.SendJSON(c, status, fiber.Map{
			"message": err.Error(),
		})
	}

	return responses.SendJSON(c, status, fiber.Map{
		"message": "success",
	})
}
