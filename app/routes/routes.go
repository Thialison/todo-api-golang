package routes

import (
	"todo-api/app/todo/controllers"
	"todo-api/app/todo/models"
	"todo-api/app/todo/services"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	model := models.NewTodoRepository()
	service := services.NewTodoService(model)
	controller := controllers.NewTodoRestController(service)
	router.GET("/", controller.Welcome)
	router.GET("/todos", controller.GetAllTodos)
	router.POST("/todo", controller.Create)
	// router.GET("/todo/:todoId", controller.GetTodo)
	// router.PUT("/todo/:todoId", controller.EditTodo)
	// router.DELETE("/todo/:todoId", controller.DeleteTodo)
	router.NoRoute(controller.NotFound)
}
