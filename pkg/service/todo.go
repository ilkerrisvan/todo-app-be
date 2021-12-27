package service

import (
	"todo-be/pkg/model"
	"todo-be/pkg/repository"
)

type TodoService struct {
	TodoRepository *repository.Repository
}
func NewTodoService(p *repository.Repository) TodoService {
	return TodoService{TodoRepository: p}
}
func (p *TodoService) AllTasks() ([]model.Task, error) {
	return p.TodoRepository.AllTasks()
}
func (p *TodoService) CreateNewTask(task *model.Task) (*model.Task, error) {
	return p.TodoRepository.CreateNewTask(task)
}


