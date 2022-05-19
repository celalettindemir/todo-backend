package repository

import (
	"backend/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTodoRepository(t *testing.T) {
	t.Run("Db all todo", func(t *testing.T) {
		db = []model.ToDo{{Id: 1, Task: "Do it"}}
		todoRepository := NewTodoRepository()
		todos, _ := todoRepository.FindAllTodos()
		assert.Equal(t, todos, []model.ToDo{{Id: 1, Task: "Do it"}})
	})
}
func TestAddTodo(t *testing.T) {
	t.Run("Db add task", func(t *testing.T) {
		db = []model.ToDo{{Id: 1, Task: "Do it"}}
		todoRepository := NewTodoRepository()
		todos, _ := todoRepository.AddTodo("Drink water")
		assert.Equal(t, todos, model.ToDo{Id: 2, Task: "Drink water"})
	})
}
