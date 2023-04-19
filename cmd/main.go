package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/joho/godotenv"
	"github.com/m3rashid-org/hmis-go-server/internal/config"
	"github.com/m3rashid-org/hmis-go-server/internal/controllers"
	"github.com/m3rashid-org/hmis-go-server/internal/middlewares"
	"github.com/m3rashid-org/hmis-go-server/internal/models"
	"github.com/m3rashid-org/hmis-go-server/internal/utils"
	"github.com/m3rashid-org/hmis-go-server/internal/ws"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	client, ctx, cancel, err := models.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer models.Close(client, ctx, cancel)
	models.Ping(client, ctx)

	app := config.FiberApp
	app.Use(middlewares.Cors)
	app.Use(middlewares.Csrf)
	app.Use(middlewares.Etag)
	app.Use(middlewares.Logger)
	app.Use(middlewares.Compression)
	app.Use(middlewares.Idempotency)
	app.Use(middlewares.Limiter)
	app.Use(func(c *fiber.Ctx) error {
		utils.SetLocal(c, "db", client) // add db connection to context
		utils.SetLocal(c, "ctx", ctx)   // add context to context
		return c.Next()
	})

	app.Get("/", controllers.Base)
	app.Get("/ping", controllers.Ping)
	app.Get("/metrics", controllers.MonitorEndpoint)
	app.Get("/favicon.ico", controllers.Favicon)

	app.Post("/login", controllers.Login)
	app.Post("/register", controllers.CreateUser)

	socket := app.Group("/ws", ws.UpgraderMiddleware)
	socket.Get("/", websocket.New(ws.WebsocketHandler))
	go ws.RunHub()

	// auth group for all protected requests
	auth := app.Group("/auth", middlewares.Protected())
	auth.Post("/ping", controllers.AuthPing)
	auth.Post("/user", controllers.CurrentUser)
	auth.Post("/logout", controllers.Logout)
	auth.Post("/refresh", controllers.RefreshToken)

	app.Use(middlewares.Recover)
	log.Fatal(app.Listen(":" + os.Getenv("HTTP_PORT")))
}
