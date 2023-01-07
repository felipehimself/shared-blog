package responses

import "github.com/gofiber/fiber/v2"

func SendJSON(c *fiber.Ctx, status int, data interface{}) error {
	c.Accepts("application/json")
	c.Status(status)
	return c.JSON(data)

}

func Error(c *fiber.Ctx, status int, data fiber.Map) error {

	c.Status(status)
	return c.JSON(data)

}
