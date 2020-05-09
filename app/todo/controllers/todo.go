package controllers

import (
	"net/http"
	"todo-api/app/todo/services"

	"github.com/gin-gonic/gin"
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

func (controller *TodoController) Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to TODO API",
	})
	return
}

func (controller *TodoController) NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Route Doesn't Exist",
	})
	return
}
