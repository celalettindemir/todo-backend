package service

import (
	"backend/mock"
	"backend/model"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTodos(t *testing.T) {
	t.Run("Db all todo", func(t *testing.T) {
		mockData := []model.ToDo{
			model.ToDo{Id: 1, Task: "Do it"},
			model.ToDo{Id: 2, Task: "Drink water"},
		}
		mockService := mock.NewMockTodoRepository(gomock.NewController(t))
		mockService.EXPECT().FindAllTodos().Return(mockData, nil).Times(1)
		todoService := NewTodoService(mockService)
		todos, _ := todoService.Todos()
		assert.Equal(t, todos,
			[]model.ToDo{
				{Id: 1, Task: "Do it"},
				{Id: 2, Task: "Drink water"},
			})
	})
}
func TestCreateTodo(t *testing.T) {
	t.Run("Db add task", func(t *testing.T) {
		mockData := model.ToDo{Id: 1, Task: "Drink water"}
		mockService := mock.NewMockTodoRepository(gomock.NewController(t))
		mockService.EXPECT().AddTodo("Drink water").Return(mockData, nil).Times(1)
		todoService := NewTodoService(mockService)
		todos, _ := todoService.CreateTodo("Drink water")
		assert.Equal(t, todos,
			model.ToDo{Id: 1, Task: "Drink water"})
	})
	t.Run("Db add task problem", func(t *testing.T) {
		mockService := mock.NewMockTodoRepository(gomock.NewController(t))
		mockService.EXPECT().AddTodo("Drink water").Return(model.ToDo{}, errors.New("task could not be added")).Times(1)
		todoService := NewTodoService(mockService)
		_, err := todoService.CreateTodo("Drink water")
		assert.Equal(t, err.Error(), "task could not be added")
	})
}
