package controller

import (
	"github.com/gofiber/fiber/v2"
)

func Status(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Server is working!"})
}
