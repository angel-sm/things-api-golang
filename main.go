package main

import (
	"crud/thing"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/things", thing.GetAll)
	app.Post("/things", thing.Create)
	app.Get("/things/:id", thing.GetById)
	app.Put("/things/:id", thing.Update)
	app.Delete("/things/:id", thing.Delete)

	app.Listen(":3000")
}
