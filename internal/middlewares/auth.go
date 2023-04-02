package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET_KEY")),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"status": "error",
				// TODO: Give Proper Error Message to the client
				"message": "Missing or malformed JWT",
				"data":    nil,
			})
	}

	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{
			"status": "error",
			// TODO: Give Proper Error Message to the client
			"message": "Invalid or expired JWT",
			"data":    nil,
		})
}
