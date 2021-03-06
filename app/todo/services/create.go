package services

import (
	"todo-api/app/todo/models"
	"todo-api/app/todo/repositories"
	"todo-api/app/todo/views/requests"
)

func (service *TodoService) Create(
	request requests.Create,
) (*models.Todo, error) {

	todo := repositories.NewTodo(
		request.Title,
		request.Body,
		request.Status,
	)

	model, err := service.Repo.Create(todo)
	if err != nil {
		return model, err
	}

	return model, nil

}
