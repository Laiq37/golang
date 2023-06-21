package middleware

import (
	"ecommerce-store-service/auth"

	"github.com/gofiber/fiber/v2"
)

type ValidUser struct {
	Email string
	Type  string
}

func VerifyUser(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "authorization required!",
			"data":    nil,
		})
	}
	user, err := auth.Validate(token)
	if err != nil {
		return c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	c.Locals("user", user)
	return c.Next()
}
