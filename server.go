package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"Go_Fiber_Starter/database"
	"Go_Fiber_Starter/router"
)

func main() {
	app := fiber.New()
	
	database.Connect()
	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}