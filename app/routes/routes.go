package routes

import (
	"todo-api/app/todo/controllers"
	"todo-api/app/todo/repositories"
	"todo-api/app/todo/services"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	repository := repositories.NewTodoRepo()
	service := services.NewTodoService(repository)
	controller := controllers.NewTodoRestController(service)
	router.GET("/", controller.Welcome)
	router.GET("/todos", controller.GetAllTodos)
	router.POST("/todo", controller.Create)
	router.GET("/todos/:title", controller.GetTodoByTitle)
	// router.PUT("/todo/:todoId", controller.EditTodo)
	// router.DELETE("/todo/:todoId", controller.DeleteTodo)
	router.NoRoute(controller.NotFound)
}
