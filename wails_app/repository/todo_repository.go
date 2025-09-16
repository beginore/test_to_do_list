package repositories

import (
	"database/sql"
	"time"
	"wails_app/models"
)

type TodoRepository interface {
	GetAll() ([]models.Todo, error)
	Create(todo *models.Todo) error
	Toggle(id int) error
	Delete(id int) error
	DeleteAll() error
}

type todoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) GetAll() ([]models.Todo, error) {
	query := `SELECT id, title, completed, created_at, completed_at FROM todos ORDER BY  created_at DESC `
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		var completedAt sql.NullTime

		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.CompletedAt)
		if err != nil {
			return nil, err
		}
		if completedAt.Valid {
			todo.CompletedAt = &completedAt.Time
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *todoRepository) Create(todo *models.Todo) error {
	query := `INSERT INTO todos (title, completed, created_at) VALUES ($1, $2, $3) RETURNING id`
	return r.db.QueryRow(query, todo.Title, todo.Completed, time.Now()).Scan(&todo.ID)
}

func (r *todoRepository) Toggle(id int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	var completed bool
	var completedAt sql.NullTime
	err = tx.QueryRow("SELECT completed, completed_at FROM todos WHERE id = $1", id).Scan(&completed, &completedAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	newCompleted := !completed
	var newCompletedAt interface{}
	if newCompleted {
		newCompletedAt = time.Now()
	} else {
		newCompletedAt = nil
	}

	_, err = tx.Exec("UPDATE todos SET completed = $1, completed_at = $2 WHERE id = $3",
		newCompleted, newCompletedAt, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *todoRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM todos WHERE id = $1", id)
	return err
}

func (r *todoRepository) DeleteAll() error {
	_, err := r.db.Exec("DELETE FROM todos")
	return err
}
