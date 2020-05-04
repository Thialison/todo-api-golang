package models

import (
	"fmt"
	"log"
	"net/http"
	"todo-api/go_db/config"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Body      string `json:"body"`
	Completed string `json:"completed"`
}

func CreateTodoTable(db *gorm.DB) error {
	db.DropTableIfExists("todos")
	db.Debug().AutoMigrate(&Todo{})
	log.Printf("Todo table created")
	return nil
}

func Create(c *gin.Context) {
	var request Todo
	db := config.GetDB()

	c.ShouldBindJSON(&request)

	fmt.Println(request)

	todo := Todo{Title: request.Title, Body: request.Body, Completed: request.Completed}

	db.NewRecord(todo)

	err := db.Create(&todo).Error

	if err != nil {
		db.Rollback()
		fmt.Println("Models: ", err)
	}

	db.NewRecord(todo)

	c.JSON(http.StatusOK, gin.H{
		"message": todo,
	})
	return
}
