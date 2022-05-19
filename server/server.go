package server

import (
	"backend/controllers"
	"backend/repository"
	"backend/service"
	"fmt"
	"net/http"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) StartServer(port int) error {
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository)
	todoController := controllers.NewTodoController(todoService)
	routes := []routePack{
		newRoutePack("GET", "/api", todoController.GetAllTodos),
		newRoutePack("POST", "/api", todoController.PostTodos),
	}
	router := NewRoute(routes)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), http.HandlerFunc(router.Serve))
	return err
}
