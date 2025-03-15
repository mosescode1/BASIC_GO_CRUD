package service

import (
	"context"

	"github.com/mosescode1/BASIC_GO_CRUD.git/internals/entity"
	"github.com/mosescode1/BASIC_GO_CRUD.git/internals/repo"
)




type TodoService struct {
	ctx context.Context
}

func NewTodoService(ctx context.Context) *TodoService {
	return &TodoService{
		ctx: ctx,
	}
}

func (s *TodoService) CreateTodo(todo *entity.Todo) (*entity.Todo, error) {
	result, err := repo.NewTodoRepo(s.ctx).CreateTodo(todo)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *TodoService) GetTodo(id string) (*entity.Todo, error) {

	todo, err := repo.NewTodoRepo(s.ctx).GetTodo(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *TodoService) GetAllTodos() ([]*entity.Todo, error) {
	todos, err := repo.NewTodoRepo(s.ctx).GetAllTodos()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *TodoService) UpdateTodo(id string, todo *entity.Todo) (*entity.Todo, error) {

	result, err := repo.NewTodoRepo(s.ctx).UpdateTodo(id, todo)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *TodoService) DeleteTodo(id string) error {
	err := repo.NewTodoRepo(s.ctx).DeleteTodo(id)
	if err != nil {
		return err
	}
	return nil
}