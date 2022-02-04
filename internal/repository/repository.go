package repository

import (
	"database/sql"
	"todolist/m/internal/models"
)

type TaskRepository interface {
	AddTask(value models.Task) (int, error)
	GetById(id int) (*models.Task, error)
}

type taskRepository struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{
		DB: db,
	}
}

func (t *taskRepository) AddTask(value models.Task) (int, error) {

	query := `INSERT INTO tasks("title", "description", "done") VALUES ($1, $2, $3) RETURNING id`

	var id int
	err := t.DB.QueryRow(query, value.Title, value.Description, value.Done).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (t *taskRepository) GetById(id int) (*models.Task, error) {

	query := `SELECT title, description, done FROM tasks WHERE id = $1`

	var task models.Task
	err := t.DB.QueryRow(query, id).Scan(&task.Title, &task.Description, &task.Done)
	if err != nil {
		return nil, err
	}

	return &task, nil
}
