package controllers

import (
	"GoProject/config"
	"GoProject/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	config.DB.Find(&users)

	return c.JSON(users)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	type UpdateRequest struct {
		Role       string `json:"role"`
		Statusdata string `json:"statusdata"`
	}

	var req UpdateRequest
	c.BodyParser(&req)

	user.Role = req.Role
	user.Statusdata = req.Statusdata

	config.DB.Save(&user)

	return c.JSON(fiber.Map{
		"message": "User updated",
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to delete user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User deleted",
	})
}
