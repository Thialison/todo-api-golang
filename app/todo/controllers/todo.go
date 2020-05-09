package controllers

import (
	"todo-api/app/todo/services"
)

type TodoController struct {
	Service *services.TodoService
}

func NewTodoRestController(
	service *services.TodoService,
) *TodoController {
	return &TodoController{
		Service: service,
	}
}
