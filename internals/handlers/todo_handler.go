package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/mosescode1/BASIC_GO_CRUD.git/internals/entity"
	"github.com/mosescode1/BASIC_GO_CRUD.git/internals/service"
	"github.com/mosescode1/BASIC_GO_CRUD.git/pkg/response"
)

type TodoHandler struct {
	ctx context.Context
}

func NewTodoHandler(ctx context.Context) *TodoHandler {
	return &TodoHandler{
		ctx: ctx,
	}
}

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var input entity.Todo
	if err := c.BodyParser(&input); err != nil {
		return response.UnprocessableEntity(c, "Invalid request payload", err)
	}
	if input.Title == "" {
		return response.UnprocessableEntity(c, "Title is required", nil)
	}
	if input.Description == "" {
		return response.UnprocessableEntity(c, "Description is required", nil)
	}

	todo, err := service.NewTodoService(h.ctx).CreateTodo(&input)

	if err != nil {
		return response.InternalServerError(c, err, "Failed to create todo")
	}

	todoMap := map[string]interface{}{
		"id":          todo.ID,
		"title":       todo.Title,
		"description": todo.Description,
		"created_at":  todo.CreatedAt,
		"updated_at":  todo.UpdatedAt,
	}
	return response.Success(c,"Todo created successfully", todoMap)
}


func (h *TodoHandler) GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	todo, err := service.NewTodoService(h.ctx).GetTodo(id)
	if err != nil {
		return response.NotFound(c,"Todo not found")
	}
	todoMap := map[string]interface{}{
		"id":          todo.ID,
		"title":       todo.Title,
		"description": todo.Description,
		"created_at":  todo.CreatedAt,
		"updated_at":  todo.UpdatedAt,
	}
	return response.Success(c,"Todo retrieved successfully", todoMap)
}

func (h *TodoHandler) GetAllTodos(c *fiber.Ctx) error {
	todos, err := service.NewTodoService(h.ctx).GetAllTodos()
	if err != nil {
		return response.InternalServerError(c, err, "Failed to retrieve todos")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Todos retrieved successfully",
		"data":    todos,
	})
}

func (h *TodoHandler) UpdatedTodo(c *fiber.Ctx) error {
	return response.Success(c,"Todo updated successfully", nil)
}

func (h *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	return response.Success(c,"Todo deleted successfully", nil)
}