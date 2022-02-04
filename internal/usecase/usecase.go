package usecase

import (
	"errors"
	"todolist/m/internal/models"
	"todolist/m/internal/repository"
)

type TaskUsecase interface {
	AddTask(task models.Task) (models.Task, error)
	DeleteTask(id int) error
	List() ([]models.Task, error)
	Mark(id int, request models.MarkRequest) (models.Task, error)
}

type taskUsecase struct {
	repo repository.TaskRepository
}

func NewTaskUsecase(repo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{repo: repo}
}

func (u *taskUsecase) AddTask(task models.Task) (models.Task, error) {
	if task.Title == "" {
		return models.Task{}, errors.New("empty title")
	}

	id, err := u.repo.AddTask(task)
	if err != nil {
		return models.Task{}, err
	}

	result, err := u.repo.GetById(id)
	if err != nil {
		return models.Task{}, err
	}

	return result, nil
}

func (u *taskUsecase) DeleteTask(id int) error {
	return u.repo.DeleteTask(id)
}

func (u *taskUsecase) List() ([]models.Task, error) {
	return u.repo.GetAll()
}

func (u *taskUsecase) Mark(id int, request models.MarkRequest) (models.Task, error) {
	err := u.repo.UpdateDone(id, request.Done)
	if err != nil {
		return models.Task{}, err
	}

	result, err := u.repo.GetById(id)
	if err != nil {
		return models.Task{}, err
	}

	return result, nil
}
