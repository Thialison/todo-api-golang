package repositories

import (
	"fmt"
	"todo-api/app/todo/models"
	"todo-api/go_db/config"
)

type TodoRepo struct {
}

func NewTodoRepo() *TodoRepo {
	return &TodoRepo{}
}

func NewTodo(
	title string,
	body string,
	status string,
) *models.Todo {
	return &models.Todo{
		Title:  title,
		Body:   body,
		Status: status,
	}
}

func (repository *TodoRepo) Create(
	todo *models.Todo,
) (*models.Todo, error) {
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

func (repository *TodoRepo) GetAll() ([]models.Todo, error) {
	db := config.GetDB()
	var todos []models.Todo
	err := db.Select([]string{"id, title, body, status"}).Find(&todos).Error
	return todos, err
}

func (repository *TodoRepo) GetTodoByTitle(
	title string,
) ([]models.Todo, error) {
	db := config.GetDB()

	var todos []models.Todo
	err := db.Where("title = ?", title).Find(&todos).Error

	return todos, err
}
