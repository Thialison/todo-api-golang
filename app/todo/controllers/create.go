package controllers

import (
	"fmt"
	"net/http"
	"todo-api/app/todo/views/requests"

	"github.com/gin-gonic/gin"
)

func (controller *TodoController) Create(
	c *gin.Context,
) {
	var request requests.Create
	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(request)

	responseBody, err := controller.Service.Create(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": responseBody})
	return
}
