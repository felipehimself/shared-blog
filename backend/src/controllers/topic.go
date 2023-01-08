package controllers

import (
	"shared-blog-backend/src/database"
	"shared-blog-backend/src/repositories"
	"shared-blog-backend/src/responses"

	"github.com/gofiber/fiber/v2"
)

func GetTopics(c *fiber.Ctx) error {

	db, err := database.Connect()

	if err != nil {
		return responses.Error(c, fiber.StatusBadRequest, fiber.Map{
			"message": err.Error(),
		})
	}

	defer db.Close()

	repo := repositories.TopicRepository(db)

	topics, err := repo.Topics()

	if err != nil {
		return responses.Error(c, fiber.StatusBadRequest, fiber.Map{
			"message": err.Error(),
		})
	}

	return responses.SendJSON(c, fiber.StatusOK, topics)
}
