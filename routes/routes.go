package routes

import (
	"example.com/jotbytes-server/handlers"
	"example.com/jotbytes-server/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/notes", handlers.CreateNote)
	app.Get("/api/notes", handlers.GetNotes)
	app.Post("/api/secure/notes", middleware.BasicAuth(), handlers.CreateSecureNote)
}
