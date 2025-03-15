package repo

import (
	"context"
	"errors"

	"github.com/mosescode1/BASIC_GO_CRUD.git/internals/database"
	"github.com/mosescode1/BASIC_GO_CRUD.git/internals/entity"
	"github.com/mosescode1/BASIC_GO_CRUD.git/pkg"
	"gorm.io/gorm"
)

type TodoRepo struct {
	ctx context.Context
}


func NewTodoRepo(ctx context.Context) *TodoRepo {
	return &TodoRepo{
		ctx: ctx,
	}
}

func (r *TodoRepo) CreateTodo(todo *entity.Todo) (*entity.Todo ,error) {
	result := database.DB.Create(&todo)
	if result.Error != nil {
		return nil, result.Error
	}
	return todo, nil
}

func (r *TodoRepo) GetTodo(id string) (*entity.Todo, error) {
	var todo *entity.Todo
	result := database.DB.Where("id = ?", id).First(&todo)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, pkg.ErrRecordNotFound
	} else if result.Error != nil {
		return nil, result.Error
	}

	return todo, nil

}

func (r *TodoRepo) GetAllTodos() ([]*entity.Todo, error) {
	var todos []*entity.Todo
	result := database.DB.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}


func (r *TodoRepo) UpdateTodo(id string, todo *entity.Todo) (*entity.Todo, error) {
	var input *entity.Todo

	result := database.DB.Model(&entity.Todo{}).Where("id = ?", id).Updates(input)
	if result.Error != nil {
		return nil, result.Error
	}
	return input, nil
}

func (r *TodoRepo) DeleteTodo(id string) error {
	result := database.DB.Where("id = ?", id).Delete(&entity.Todo{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}