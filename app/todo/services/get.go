package services

import "todo-api/app/todo/models"

func (service *TodoService) GetAll() ([]models.Todo, error) {
	todos, err := service.Repo.GetAll()
	return todos, err
}

func (service *TodoService) GetTodoByTitle(
	title string,
) ([]models.Todo, error) {
	todos, err := service.Repo.GetTodoByTitle(title)
	return todos, err
}
