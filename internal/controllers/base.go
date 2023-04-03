package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Base(c *fiber.Ctx) error {
	return c.SendString("Hello from HMIS")
}

func Ping(c *fiber.Ctx) error {
	return c.SendString("PONG !")
}

var MonitorEndpoint = monitor.New(monitor.Config{
	Title:      "HMIS Server Metrics",
	CustomHead: "HMIS",
})
