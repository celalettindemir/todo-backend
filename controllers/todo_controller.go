package controllers

import (
	"backend/service"
	"encoding/json"
	"io"
	"net/http"
)

type TodoController interface {
	GetAllTodos(http.ResponseWriter, *http.Request)
	PostTodos(http.ResponseWriter, *http.Request)
}
type todoController struct {
	s service.ToDoService
}

func (c *todoController) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	todos, err := c.s.Todos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	todosBytes, err := json.Marshal(todos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(todosBytes)
}

type PostTodoDTO struct {
	Task string `json:"task"`
}

func (c *todoController) PostTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	defer r.Body.Close()
	body, _ := io.ReadAll(r.Body)

	var postTodoDTO PostTodoDTO
	err := json.Unmarshal(body, &postTodoDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	todos, err := c.s.CreateTodo(postTodoDTO.Task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	todosBytes, _ := json.Marshal(todos)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(todosBytes)
}
func NewTodoController(s service.ToDoService) TodoController {
	return &todoController{s}
}
