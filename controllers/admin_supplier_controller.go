package controllers

import (
	"GoProject/config"
	"GoProject/models"

	"github.com/gofiber/fiber/v2"
)

func GetSuppliers(c *fiber.Ctx) error {
	var suppliers []models.Supplier
	config.DB.Find(&suppliers)

	return c.JSON(suppliers)
}

func AddSupplier(c *fiber.Ctx) error {
	
	var suppliers models.Supplier
	
	type UpdateRequest struct {
		Nama       	string `json:"nama"`
		Email       string `json:"email"`
		Alamat 		string `json:"alamat"`
	}

	var req UpdateRequest
	c.BodyParser(&req)

	suppliers.Nama = req.Nama
	suppliers.Email = req.Email
	suppliers.Alamat = req.Alamat

	config.DB.Create(&suppliers)

	return c.JSON(fiber.Map{
		"message": "Data success added",
	})
}

func UpdateSupplier(c *fiber.Ctx) error {
	id := c.Params("id")
	var supplier models.Supplier

	if err := config.DB.First(&supplier, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Supplier not found",
		})
	}

	type UpdateRequest struct {
		Nama       	string `json:"nama"`
		Email       string `json:"email"`
		Alamat 		string `json:"alamat"`
	}

	var req UpdateRequest
	c.BodyParser(&req)

	supplier.Nama = req.Nama
	supplier.Email = req.Email
	supplier.Alamat = req.Alamat

	config.DB.Save(&supplier)

	return c.JSON(fiber.Map{
		"message": "Supplier updated",
	})
}

func DeleteSupplier(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := config.DB.Delete(&models.Supplier{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to delete user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Supplier deleted",
	})
}
