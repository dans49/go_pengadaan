package controllers

import "github.com/gofiber/fiber/v2"

func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "API jalan",
	})
}

func Nama(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"nama": "Wildan N",
	})
}
