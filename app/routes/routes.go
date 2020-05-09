package routes

import (
	"net/http"

	"todo-api/app/todo/controllers"
	"todo-api/app/todo/models"
	"todo-api/app/todo/services"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	model := models.NewTodoRepository()
	service := services.NewTodoService(model)
	controller := controllers.NewTodoRestController(service)
	router.GET("/", welcome)
	// router.GET("/todos", controllers.GetAllTodos)
	router.POST("/todo", controller.Create)
	// router.GET("/todo/:todoId", controllers.GetSingleTodo)
	// router.PUT("/todo/:todoId", controllers.EditTodo)
	// router.DELETE("/todo/:todoId", controllers.DeleteTodo)
	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome To API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Route Doesn't Exist",
	})
	return
}
