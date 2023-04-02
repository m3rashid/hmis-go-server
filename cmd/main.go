package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/m3rashid-org/hmis-go-server/internal/config"
	"github.com/m3rashid-org/hmis-go-server/internal/controllers"
	"github.com/m3rashid-org/hmis-go-server/internal/middlewares"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	app := config.FiberApp
	app.Use(middlewares.Cors)
	app.Use(middlewares.Csrf)
	app.Use(middlewares.Etag)
	app.Use(middlewares.Logger)
	app.Use(middlewares.Compression)
	app.Use(middlewares.Idempotency)
	app.Use(middlewares.Limiter)

	app.Get("/metrics", controllers.MonitorEndpoint)
	/**
	* Other routes
	 */

	app.Use(middlewares.Recover)
	log.Fatal(app.Listen(":" + os.Getenv("HTTP_PORT")))
}
