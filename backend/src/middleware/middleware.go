package middleware

import (
	"net/http"
	"shared-blog-backend/src/responses"
	"shared-blog-backend/src/utils"

	"github.com/gofiber/fiber/v2"
)

func ProtectedRoute() fiber.Handler {

	return func(c *fiber.Ctx) error {

		if _, err := utils.IsTokenValid(c); err != nil {
			return responses.Error(c, http.StatusForbidden, fiber.Map{
				"message": err.Error(),
			})
		}

		return c.Next()
	}
}
