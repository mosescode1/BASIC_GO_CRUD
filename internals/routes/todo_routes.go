package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/mosescode1/BASIC_GO_CRUD.git/internals/handlers"
)


func SetUpTodoRoutes(app *fiber.App) {

	todo := app.Group("/todos")
	todoHandler := handlers.NewTodoHandler(context.Background())
	todo.Post("/", todoHandler.CreateTodo)
	todo.Get("/", todoHandler.GetAllTodos)
	todo.Get("/:id", todoHandler.GetTodo)
	todo.Patch("/:id", todoHandler.UpdatedTodo)
	todo.Delete("/:id", todoHandler.DeleteTodo)
}

