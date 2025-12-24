package middleware

import "github.com/gofiber/fiber/v2"

func AdminOnly(c *fiber.Ctx) error {
	role := c.Locals("role")

	if role != "ADMIN" {
		return c.Status(403).JSON(fiber.Map{
			"message": "Access denied. Admin only.",
		})
	}

	return c.Next()
}
