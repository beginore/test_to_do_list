package services

import (
	"wails_app/models"
	repositories "wails_app/repository"
)

type TodoService interface {
	GetAllTodos() ([]models.Todo, error)
	AddTodo(title string) (*models.Todo, error)
	ToggleTodo(id int) error
	DeleteTodo(id int) error
	DeleteAllTodos() error
}

type todoService struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) GetAllTodos() ([]models.Todo, error) {
	return s.repo.GetAll()
}

func (s *todoService) AddTodo(title string) (*models.Todo, error) {
	todo := &models.Todo{
		Title:     title,
		Completed: false,
	}

	err := s.repo.Create(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *todoService) ToggleTodo(id int) error {
	return s.repo.Toggle(id)
}

func (s *todoService) DeleteTodo(id int) error {
	return s.repo.Delete(id)
}

func (s *todoService) DeleteAllTodos() error {
	return s.repo.DeleteAll()
}
