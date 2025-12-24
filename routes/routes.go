package routes

import (
	"GoProject/middleware"
	"GoProject/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Fiber jalan 🔥")
	// })

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./public/login.html")
	})

	app.Get("/car2", func(c *fiber.Ctx) error {
		return c.SendString("12")
	})

	api.Get("/hello", controllers.Hello)
	api.Get("/nama", controllers.Nama)

	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)


	api.Get("/profile", middleware.JWTProtected, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"user_id":  c.Locals("user_id"),
			"username": c.Locals("username"),
			"nama":     c.Locals("nama"),
			"role":     c.Locals("role"),
		})
	})
	app.Get(
		"/admin/dashboard",
		middleware.JWTProtected,
		middleware.AdminOnly,
		func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"message": "Welcome Admin",
			})
		},
	)

	admin := api.Group("/admin", middleware.JWTProtected, middleware.AdminOnly)

	admin.Get("/users", controllers.GetUsers)
	admin.Put("/users/:id", controllers.UpdateUser)
	admin.Delete("/users/:id", controllers.DeleteUser)


}
