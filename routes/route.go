package routes

import (
	"golang-simple-crud/controller"

	"github.com/gofiber/fiber/v2"
)

func ApiRoute(app *fiber.App) {
	route := app.Group("/api/todos")
	route.Get("/", controller.GetTodos)
	route.Get("/:id", controller.GetTodo)
	route.Post("/", controller.CreateTodo)
	route.Patch("/:id", controller.UpdateTodo)
	route.Delete("/:id", controller.DeleteTodo)
}
