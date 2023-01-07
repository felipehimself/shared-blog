package controllers

import (
	"shared-blog-backend/src/config"
	"shared-blog-backend/src/database"
	"shared-blog-backend/src/models"
	"shared-blog-backend/src/repositories"
	"shared-blog-backend/src/responses"
	"shared-blog-backend/src/utils"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	fiber "github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return responses.Error(c, fiber.StatusUnprocessableEntity, fiber.Map{
			"message": "error while reading body",
		})

	}

	if err := user.ValidateFields(); err != nil {
		return responses.Error(c, fiber.StatusUnprocessableEntity, fiber.Map{
			"message": err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		return responses.Error(c, fiber.StatusUnprocessableEntity, fiber.Map{
			"message": "error when creating account",
		})
	}

	user.Password = string(hashedPassword)

	db, err := database.Connect()

	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	defer db.Close()

	repo := repositories.UserRepository(db)

	if err = repo.SignUpUser(user); err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	return responses.SendJSON(c, fiber.StatusCreated, fiber.Map{
		"message": "user created",
	})

}

func SignIn(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return responses.Error(c, fiber.StatusUnprocessableEntity, fiber.Map{
			"message": "error while reading body",
		})
	}

	db, err := database.Connect()

	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	defer db.Close()

	repo := repositories.UserRepository(db)

	userId, err := repo.SignInUser(user)

	if err != nil {
		return responses.Error(c, fiber.StatusBadRequest, fiber.Map{
			"message": err.Error(),
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(userId)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(config.SECRET_KEY))

	if err != nil {
		return responses.Error(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	cookie := fiber.Cookie{
		Name:     "shared-blog-jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return responses.SendJSON(c, fiber.StatusAccepted, fiber.Map{
		"success": true,
	})
}

func SignOut(c *fiber.Ctx) error {

	cookie := fiber.Cookie{
		Name:     "shared-blog-jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return responses.SendJSON(c, fiber.StatusAccepted, fiber.Map{
		"success": true,
	})

}

func IsAuthorized(c *fiber.Ctx) error {

	if _, err := utils.IsTokenValid(c); err != nil {
		return responses.Error(c, fiber.StatusForbidden, fiber.Map{
			"success": false,
		})
	}

	return responses.SendJSON(c, fiber.StatusOK, fiber.Map{
		"success": true,
	})

}
