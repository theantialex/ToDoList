package repository

import (
	"database/sql"
	"log"
	"todolist/m/internal/models"
)

type TaskRepository interface {
	AddTask(task models.Task) (int, error)
	GetById(id int) (models.Task, error)
	DeleteTask(id int) error
	GetAll() ([]models.Task, error)
	UpdateDone(id int, done bool) error
}

type taskRepository struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{
		DB: db,
	}
}

func (r *taskRepository) AddTask(task models.Task) (int, error) {
	query := `INSERT INTO tasks("title", "description", "done") VALUES ($1, $2, $3) RETURNING id`

	var id int
	err := r.DB.QueryRow(query, task.Title, task.Description, task.Done).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *taskRepository) GetById(id int) (models.Task, error) {
	query := `SELECT id, title, description, done FROM tasks WHERE id = $1`

	var task models.Task
	err := r.DB.QueryRow(query, id).Scan(&task.Id, &task.Title, &task.Description, &task.Done)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) DeleteTask(id int) error {
	query := `DELETE FROM Tasks WHERE id = $1`
	err := r.DB.QueryRow(query, id).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) GetAll() ([]models.Task, error) {
	query := `SELECT id, title, description, done FROM tasks`

	rows, err := r.DB.Query(query)
	if err != nil {
		return []models.Task{}, err
	}

	var tasks []models.Task
	for rows.Next() {
		var task models.Task

		err = rows.Scan(&task.Id, &task.Title, &task.Description, &task.Done)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	err = rows.Close()
	if err != nil {
		log.Fatal("Error occurred during closing rows")
	}

	return tasks, nil
}

func (r *taskRepository) UpdateDone(id int, done bool) error {
	query := `UPDATE tasks SET "done" = $1 WHERE id = $2`
	err := r.DB.QueryRow(query, done, id).Err()
	if err != nil {
		return err
	}

	return nil
}
