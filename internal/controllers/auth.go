package controllers

import "github.com/gofiber/fiber/v2"

func CurrentUser (c *fiber.Ctx) error {
	return	c.SendString("Hello, World!")
}

func AuthPing (c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func Login (c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func CreateUser (c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func RegisterInitStepOne (c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func RegisterInitStepTwo (c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func BlockUser (c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

