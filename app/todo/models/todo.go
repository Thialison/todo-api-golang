package models

import (
	"fmt"
	"log"
	"todo-api/go_db/config"

	"github.com/jinzhu/gorm"
)

type TodoRepo struct {
}

func NewTodoRepository() *TodoRepo {
	return &TodoRepo{}
}

type Todo struct {
	gorm.Model
	Title  string `gorm:"column:title;not null"`
	Body   string `gorm:"column:body"`
	Status string `gorm:"column:status;not null"`
}

func NewCommission(
	title string,
	body string,
	status string,
) *Todo {
	return &Todo{
		Title:  title,
		Body:   body,
		Status: status,
	}
}

func (repository *TodoRepo) Create(
	todo *Todo,
) (*Todo, error) {
	db := config.GetDB()

	db.NewRecord(todo)

	err := db.Create(todo).Error
	if err != nil {
		db.Rollback()
		fmt.Println("Models: ", err)
	}

	db.NewRecord(todo)
	return todo, nil
}

func CreateTodoTable(db *gorm.DB) error {
	db.DropTableIfExists("todos")
	db.Debug().AutoMigrate(&Todo{})
	log.Printf("Todo table created")
	return nil
}
