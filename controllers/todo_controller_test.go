package controllers

import (
	"backend/mock"
	"backend/model"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAllTodos(t *testing.T) {

	t.Run("only GET method allowed", func(t *testing.T) {

		mockService := mock.NewMockToDoService(gomock.NewController(t))
		mockService.EXPECT().Todos().Return([]model.ToDo{
			{Id: 1, Task: "Do it"},
		}, nil).Times(1)
		todoController := NewTodoController(mockService)
		req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
		res := httptest.NewRecorder()
		todoController.GetAllTodos(res, req)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})
	t.Run("POST not allowed", func(t *testing.T) {
		todoController := NewTodoController(nil)

		req := httptest.NewRequest(http.MethodPost, "/", http.NoBody)
		res := httptest.NewRecorder()

		todoController.GetAllTodos(res, req)
		assert.Equal(t, http.StatusNotImplemented, res.Result().StatusCode)
	})
	t.Run("return verify data", func(t *testing.T) {
		mockService := mock.NewMockToDoService(gomock.NewController(t))
		mockService.EXPECT().Todos().Return([]model.ToDo{
			{Id: 1, Task: "Do it"},
			{Id: 2, Task: "Drink water"},
		}, nil).Times(1)
		todoController := NewTodoController(mockService)
		req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
		res := httptest.NewRecorder()
		todoController.GetAllTodos(res, req)

		var expectedResBody []model.ToDo
		err := json.Unmarshal(res.Body.Bytes(), &expectedResBody)
		assert.Nil(t, err, "json unmarshal err")
		assert.Equal(t, "application/json; charset=utf-8", res.Result().Header.Get("Content-Type"))
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
		assert.Equal(t, expectedResBody, []model.ToDo{
			{Id: 1, Task: "Do it"},
			{Id: 2, Task: "Drink water"},
		})
	})
}

func TestPostTodos(t *testing.T) {
	t.Run("Get not allowed", func(t *testing.T) {
		todoController := NewTodoController(nil)

		req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
		res := httptest.NewRecorder()

		todoController.PostTodos(res, req)
		assert.Equal(t, http.StatusNotImplemented, res.Result().StatusCode)
	})
	t.Run("wrong body for save task", func(t *testing.T) {
		mockService := mock.NewMockToDoService(gomock.NewController(t))
		todoController := NewTodoController(mockService)
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{task\"Drink water\"}"))
		res := httptest.NewRecorder()
		todoController.PostTodos(res, req)
		assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
	})
	t.Run("service bad response", func(t *testing.T) {

		mockService := mock.NewMockToDoService(gomock.NewController(t))
		todoController := NewTodoController(mockService)
		mockService.EXPECT().CreateTodo("Drink water").Return(model.ToDo{}, errors.New("Failed Save")).Times(1)
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{\"task\":\"Drink water\"}"))
		res := httptest.NewRecorder()
		todoController.PostTodos(res, req)
		assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
	})
	t.Run("service response failed convert", func(t *testing.T) {
		mockService := mock.NewMockToDoService(gomock.NewController(t))
		mockService.EXPECT().CreateTodo("Drink water").Return(model.ToDo{Id: 1, Task: "Drink water"}, nil).Times(1)
		todoController := NewTodoController(mockService)
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{\"task\":\"Drink water\"}"))
		res := httptest.NewRecorder()
		todoController.PostTodos(res, req)

		expectedResBody := model.ToDo{}
		err := json.Unmarshal(res.Body.Bytes(), &expectedResBody)
		assert.Nil(t, err, "json unmarshal err")
		assert.Equal(t, "application/json; charset=utf-8", res.Result().Header.Get("content-type"))
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
		assert.Equal(t, expectedResBody.Id, 1)
		assert.Equal(t, expectedResBody.Task, "Drink water")
	})

}
