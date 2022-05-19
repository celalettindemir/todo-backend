package server

import (
	"backend/controllers"
	"backend/repository"
	"backend/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServe(t *testing.T) {
	tests := []struct {
		method     string
		path       string
		status     int
		body       string
		expectBody string
	}{
		{"GET", "/", 200, "", "[{\"id\":1,\"task\":\"Buy car\"}]"},
		{"POST", "/", 200, "{\"task\":\"Do it\"}", "{\"id\":2,\"task\":\"Do it\"}"},
		{"GET", "/test", 404, "", ""},
		{"PUT", "/", 405, "", ""},
	}
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository)
	todoController := controllers.NewTodoController(todoService)
	routes := []routePack{
		newRoutePack("GET", "/", todoController.GetAllTodos),
		newRoutePack("POST", "/", todoController.PostTodos),
	}
	router := http.HandlerFunc(NewRoute(routes).Serve)

	t.Run("Retable", func(t *testing.T) {
		for _, test := range tests {
			path := strings.ReplaceAll(test.path, "/", "_")
			t.Run(test.method+path, func(t *testing.T) {
				recorder := httptest.NewRecorder()
				request, err := http.NewRequest(test.method, test.path, strings.NewReader(test.body))
				if err != nil {
					t.Fatal(err)
				}
				router.ServeHTTP(recorder, request)
				if recorder.Code != test.status {
					t.Fatalf("expected status %d, got %d", test.status, recorder.Code)
				}
				if test.status == 200 {
					body := recorder.Body.String()
					if body != test.expectBody {
						t.Fatalf("expected body %q, got %q", test.expectBody, body)
					}
				}
			})
		}
	})
}
