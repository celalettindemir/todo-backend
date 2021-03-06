// Code generated by MockGen. DO NOT EDIT.
// Source: .\controllers\todo_controller.go

// Package mock is a generated GoMock package.
package mock

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTodoController is a mock of TodoController interface.
type MockTodoController struct {
	ctrl     *gomock.Controller
	recorder *MockTodoControllerMockRecorder
}

// MockTodoControllerMockRecorder is the mock recorder for MockTodoController.
type MockTodoControllerMockRecorder struct {
	mock *MockTodoController
}

// NewMockTodoController creates a new mock instance.
func NewMockTodoController(ctrl *gomock.Controller) *MockTodoController {
	mock := &MockTodoController{ctrl: ctrl}
	mock.recorder = &MockTodoControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoController) EXPECT() *MockTodoControllerMockRecorder {
	return m.recorder
}

// GetAllTodos mocks base method.
func (m *MockTodoController) GetAllTodos(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetAllTodos", arg0, arg1)
}

// GetAllTodos indicates an expected call of GetAllTodos.
func (mr *MockTodoControllerMockRecorder) GetAllTodos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTodos", reflect.TypeOf((*MockTodoController)(nil).GetAllTodos), arg0, arg1)
}

// PostTodos mocks base method.
func (m *MockTodoController) PostTodos(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PostTodos", arg0, arg1)
}

// PostTodos indicates an expected call of PostTodos.
func (mr *MockTodoControllerMockRecorder) PostTodos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostTodos", reflect.TypeOf((*MockTodoController)(nil).PostTodos), arg0, arg1)
}
