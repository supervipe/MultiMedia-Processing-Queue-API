package routes

import (
	"POC-Load-Videos/handlers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/image-process", handlers.UploadImage)
}
