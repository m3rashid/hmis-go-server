package config

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

var FiberApp = fiber.New(fiber.Config{
	BodyLimit:    1024 * 1024 * 10, // 10 MB
	ServerHeader: string(AppName),
	ETag:         true,
	IdleTimeout:  5 * 60 * time.Second,
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
