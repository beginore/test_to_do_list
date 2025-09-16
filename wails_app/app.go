package main

import (
	"context"
	"fmt"
	"wails_app/services"
)

type App struct {
	ctx         context.Context
	todoService services.TodoService
}

func NewApp(todoService services.TodoService) *App {
	return &App{
		todoService: todoService,
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Add(title string) error {
	_, err := a.todoService.AddTodo(title)
	return err
}

func (a *App) List() ([]map[string]interface{}, error) {
	todos, err := a.todoService.GetAllTodos()
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, len(todos))
	for i, todo := range todos {
		result[i] = map[string]interface{}{
			"ID":          todo.ID,
			"Title":       todo.Title,
			"Completed":   todo.Completed,
			"CreatedAt":   todo.CreatedAt,
			"CompletedAt": todo.CompletedAt,
		}
	}

	return result, nil
}

func (a *App) Toggle(id int) error {
	return a.todoService.ToggleTodo(id)
}

func (a *App) Delete(id int) error {
	return a.todoService.DeleteTodo(id)
}

func (a *App) DeleteAll() error {
	return a.todoService.DeleteAllTodos()
}
