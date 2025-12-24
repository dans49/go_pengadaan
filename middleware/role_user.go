package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func UserOnly(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["role"] != "user" {
		return c.Status(403).JSON(fiber.Map{
			"message": "Access denied",
		})
	}

	return c.Next()
}
