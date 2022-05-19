package repository

import (
	"backend/model"
)

type TodoRepository interface {
	FindAllTodos() ([]model.ToDo, error)
	AddTodo(string) (model.ToDo, error)
}

var db = []model.ToDo{{Id: 1, Task: "Buy car"}}

type todoRepository struct {
}

func (r *todoRepository) AddTodo(task string) (model.ToDo, error) {
	newValue := model.ToDo{Id: len(db) + 1, Task: task}
	db = append(db, newValue)
	return newValue, nil
}

func (r *todoRepository) FindAllTodos() ([]model.ToDo, error) {

	return db, nil
}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}
