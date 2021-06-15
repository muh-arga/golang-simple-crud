package controller

import (
	"golang-simple-crud/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)


func GetTodos(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(model.Todos)
}

func GetTodo(ctx *fiber.Ctx) error {
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse id",
		})
	}

	for _, todo := range model.Todos{
		if todo.Id == id{
			return ctx.Status(fiber.StatusOK).JSON(todo)
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Error": "Data not found",
	})
}

func CreateTodo(ctx *fiber.Ctx) error{
	type requset struct {
		Name string `json:"name" form:"name"`
	}

	var body requset

	err := ctx.BodyParser(&body)
	if err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Cannot parse json",
		})
	}

	todo := &model.Todo{
		Id: len(model.Todos)+1,
		Name: body.Name,
		Completed: false,
	}

	model.Todos = append(model.Todos, todo)

	return ctx.Status(fiber.StatusOK).JSON(model.Todos)
}

func UpdateTodo(ctx *fiber.Ctx) error{
	type request struct{
		Name 		*string 	`json:"name" form:"name"`
		Completed 	*bool 		`json:"completed" form:"completed"`
	}

	var body request

	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Cannot parse json",
		})
	}

	paramsId := ctx.Params("id")

	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Cannot parse id",
		})
	}

	var todo *model.Todo

	for _, t := range model.Todos {
		if t.Id == id {
			todo = t
			break
		}
	}

	if todo == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error": "Data not found",
		})
	}

	if body.Name != nil {
		todo.Name = *body.Name
	}

	if body.Completed != nil {
		todo.Completed = *body.Completed
	}

	return ctx.Status(fiber.StatusOK).JSON(todo)
}

func DeleteTodo(ctx *fiber.Ctx) error {
	paramsId := ctx.Params("id")

	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"Error": "Cannot perse id",
		})
	}

	for i, todo := range model.Todos {
		if todo.Id == id {
			model.Todos = append(model.Todos[0:i], model.Todos[i+1:]...)
			return ctx.Status(fiber.StatusOK).JSON(model.Todos)
		}
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Error": "Data not found",
	})
}