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
		"/admin/dashboard",middleware.JWTProtected,middleware.AdminOnly,
		func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"message": "Welcome Admin",
			})
		},
	)


	app.Get("/admin/home",func(c *fiber.Ctx) error {
		return c.SendFile("./public/modul/dashboard.html")
	})
	app.Get("/admin/users",func(c *fiber.Ctx) error {
		return c.SendFile("./public/modul/users/index.html")
	});
	app.Get("/admin/suppliers",func(c *fiber.Ctx) error {
		return c.SendFile("./public/modul/suppliers/index.html")
	});
	app.Get("/admin/addsupplier",func(c *fiber.Ctx) error {
		return c.SendFile("./public/modul/suppliers/tambah.html")
	});

	app.Get("/admin/items",func(c *fiber.Ctx) error {
		return c.SendFile("./public/modul/items/index.html")
	});
	app.Get("/admin/additems",func(c *fiber.Ctx) error {
		return c.SendFile("./public/modul/items/tambah.html")
	});


	admin := api.Group("/admin", middleware.JWTProtected, middleware.AdminOnly)
	admin.Get("/users", controllers.GetUsers)
	admin.Put("/users/:id", controllers.UpdateUser)
	admin.Delete("/users/:id", controllers.DeleteUser)

	admin.Get("/suppliers", controllers.GetSuppliers)
	admin.Post("/suppliers/add", controllers.AddSupplier)
	admin.Put("/suppliers/:id", controllers.UpdateSupplier)
	admin.Delete("/suppliers/:id", controllers.DeleteSupplier)

	admin.Get("/items", controllers.GetItems)
	admin.Get("/items-dashboard", controllers.GetSisaItems)
	admin.Post("/item/add", controllers.AddItem)
	admin.Put("/items/:id", controllers.UpdateItem)
	admin.Delete("/items/:id", controllers.DeleteItem)

	// purchase := api.Group("/purchase", middleware.JWTProtected, middleware.UserOnly)

	// purchase.Post("/", controllers.CreatePurchase)


}
