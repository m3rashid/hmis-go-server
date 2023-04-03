package ws

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

/*
 * Websockets are internally TCP connections, in which the client and server can send and receive data
 * Inititally, it is a HTTP connection, and the client sends a request to upgrade to the WebSocket protocol
 * If using the javascript web Websocket API, it is handled automatically
 */

func UpgraderMiddleware(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		// Returns true if the client requested upgrade to the WebSocket protocol
		return c.Next()
	}
	return c.SendStatus(fiber.StatusUpgradeRequired)
}
