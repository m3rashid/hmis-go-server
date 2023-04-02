package config

import (
	"github.com/gofiber/fiber/v2"
)

var FiberApp = fiber.New(fiber.Config{
	BodyLimit:    1024 * 1024 * 10, // 10 MB
	ServerHeader: "HMIS",
	ETag:         true,
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		/*
			code := fiber.StatusInternalServerError
			Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
		*/
		return ctx.SendStatus(fiber.StatusInternalServerError)
	},
})
