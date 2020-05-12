package services

import (
	"todo-api/app/todo/repositories"
)

type TodoService struct {
	Repo *repositories.TodoRepo
}

func NewTodoService(
	repository *repositories.TodoRepo,
) *TodoService {
	return &TodoService{
		Repo: repository,
	}
}
