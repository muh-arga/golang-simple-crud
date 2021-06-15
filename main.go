package main

import (
	"golang-simple-crud/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Simple CRUD with Golang")

	})

	app.Get("/:name", func (ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, " + ctx.Params("name"))
	})

	routes.ApiRoute(app)

	app.Listen(":3000")
}
