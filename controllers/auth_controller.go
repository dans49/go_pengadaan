package controllers

import (
	"GoProject/config"
	"GoProject/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Kata_sandi string `json:"kata_sandi"`
	Nama_lengkap     string `json:"nama_lengkap"`
	Role    string `json:"role"`
	Statusdata string `json:"statusdata"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Kata_sandi string `json:"kata_sandi"`
}


func Register(c *fiber.Ctx) error {
	var req RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Kata_sandi), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}

	user := models.User{
		Username: 		req.Username,
		Nama_lengkap: 	req.Nama_lengkap,
		Kata_sandi: 	string(hashedPassword),
		Role:    		"USER", // default
		Statusdata: 	"AKTIF",
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Username already exists",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Register success",
	})
}

func Login(c *fiber.Ctx) error {
	var req LoginRequest
	var user models.User

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	// cari user berdasarkan username
	if err := config.DB.
		Where("username = ?", req.Username).
		First(&user).Error; err != nil {

		return c.Status(401).JSON(fiber.Map{
			"message": "Username or password wrong",
		})
	}

	// cek password
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Kata_sandi),
		[]byte(req.Kata_sandi),
	); err != nil {

		return c.Status(401).JSON(fiber.Map{
			"message": "Username or password wrong",
		})
	}

	// generate token
	token, err := config.GenerateToken(user.ID, user.Username, user.Role, user.Nama_lengkap) // diambil dari field
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login success",
		"token":   token,
	})
}

