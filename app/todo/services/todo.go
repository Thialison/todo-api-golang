package services

import "todo-api/app/todo/models"

type TodoService struct {
	Repo *models.TodoRepo
}

func NewTodoService(
	repository *models.TodoRepo,
) *TodoService {
	return &TodoService{
		Repo: repository,
	}
}
