package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Ping(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

var MonitorEndpoint = monitor.New(monitor.Config{
	Title:      "HMIS Server Metrics",
	CustomHead: "HMIS",
})
