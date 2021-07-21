package controller

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	"Go_Fiber_Starter/config"
	"Go_Fiber_Starter/models"
)

func Login(c *fiber.Ctx) error {
	var input models.LoginInput
	var user models.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error on login request"})
	}

	user = GetUserByUsername(c, input.Username)

	if user == (models.User{}) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Bad Credentials"})
	}

	if user.Password != input.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Bad Credentials"})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	
	t, err := token.SignedString([]byte(config.Config("TOKEN_KEY")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
