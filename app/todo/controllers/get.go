package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *TodoController) GetAllTodos(
	c *gin.Context,
) {

	response, err := controller.Service.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}
