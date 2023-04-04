package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/m3rashid-org/hmis-go-server/internal/config"
)

func Base(c *fiber.Ctx) error {
	return c.SendString("Hello from" + " " + string(config.AppName))
}

func Ping(c *fiber.Ctx) error {
	return c.SendString("PONG !")
}

func Favicon(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func MonitorEndpoint(c *fiber.Ctx) error {
	/*
		 * TODO: read from user if it needs
		 	- With/Without Header
			- As an API or as a Webpage
	*/

	return monitor.New(monitor.Config{
		Title:      string(config.AppName) + " " + "Server Metrics",
		CustomHead: string(config.AppName),
	})(c)
}
