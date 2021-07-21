package router

import (
	"Go_Fiber_Starter/controller"
	"Go_Fiber_Starter/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Use(logger.New())

	api := app.Group("/api")
	api.Get("/status", controller.Status)

	auth := app.Group("/auth")
	auth.Post("/login", controller.Login)

	users := app.Group("/user")
	users.Get("/:id", middleware.Protected(), controller.GetUserById)
	users.Get("/", middleware.Protected(), controller.GetAllUsers)
	users.Post("/", controller.CreateUser)
}
