package service

import (
	"backend/model"
	"backend/repository"
)

type ToDoService interface {
	Todos() ([]model.ToDo, error)
	CreateTodo(string) (model.ToDo, error)
}

type todoService struct {
	r repository.TodoRepository
}

func (t todoService) Todos() ([]model.ToDo, error) {
	result, err := t.r.FindAllTodos()
	if err != nil {
		return []model.ToDo{}, err
	}
	return result, nil
}

func (t todoService) CreateTodo(s string) (model.ToDo, error) {
	result, err := t.r.AddTodo(s)
	if err != nil {
		return model.ToDo{}, err
	}
	return result, nil
}

func NewTodoService(r repository.TodoRepository) ToDoService {
	return &todoService{r}
}
