package main

import (
	"GoProject/config"
	"GoProject/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "./public")
	
	config.ConnectDB()
	routes.SetupRoutes(app)

	app.Listen(":3000")

}
