package controllers

import (
	"GoProject/config"
	"GoProject/models"

	"github.com/gofiber/fiber/v2"
)

func GetItems(c *fiber.Ctx) error {
	var items []models.Item
	config.DB.Find(&items)

	return c.JSON(items)
}

func UpdateItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var items models.Item

	if err := config.DB.First(&items, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Item not found",
		})
	}

	type UpdateRequest struct {
		Nama_item  	string `json:"nama_item"`
		Stok       	string `json:"stok"`
		Harga 		string `json:"harga"`
	}

	var req UpdateRequest
	c.BodyParser(&req)

	items.Nama_item = req.Nama_item
	items.Stok = req.Stok
	items.Harga = req.Harga

	config.DB.Save(&items)

	return c.JSON(fiber.Map{
		"message": "Item updated",
	})
}

func DeleteItem(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := config.DB.Delete(&models.Item{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to delete user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Item deleted",
	})
}
