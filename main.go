package main

import (
	"context"
	"example.com/jotbytes-server/config"

	"example.com/jotbytes-server/middleware"
	"example.com/jotbytes-server/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDB()
	defer func() {
		if err := config.DB.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	app := fiber.New()
	routes.Setup(app)
	// Global middleware
	app.Use(middleware.RequestLogger())

	app.Listen(":3000")
}
