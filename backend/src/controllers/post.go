package controllers

import (
	"net/http"
	"shared-blog-backend/src/database"
	"shared-blog-backend/src/models"
	"shared-blog-backend/src/repositories"
	"shared-blog-backend/src/responses"
	"shared-blog-backend/src/utils"

	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {

	userOnToken, err := utils.IsTokenValid(c)

	if err != nil {
		return responses.Error(c, fiber.StatusForbidden, fiber.Map{
			"message": err.Error(),
		})
	}

	var post models.Post

	if err = c.BodyParser(&post); err != nil {
		return responses.Error(c, fiber.StatusUnprocessableEntity, fiber.Map{
			"message": "error while reading body",
		})
	}

	if err := post.ValidateFields(); err != nil {
		return responses.Error(c, fiber.StatusUnprocessableEntity, fiber.Map{
			"message": err.Error(),
		})
	}

	post.AuthorId = userOnToken

	db, err := database.Connect()

	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	defer db.Close()

	repo := repositories.PostRepository(db)

	if err = repo.Create(post); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, fiber.Map{
			"mssage": err.Error(),
		})
	}

	return responses.SendJSON(c, http.StatusCreated, fiber.Map{
		"message": "post created",
	})
}

func GetPosts(c *fiber.Ctx) error {

	db, err := database.Connect()

	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, fiber.Map{
			"mssage": err.Error(),
		})
	}

	repo := repositories.PostRepository(db)
	posts, err := repo.GetPosts()

	if err != nil {
		return responses.Error(c, fiber.StatusBadRequest, fiber.Map{
			"mssage": err.Error(),
		})
	}

	return responses.SendJSON(c, http.StatusOK, posts)

}
