package usecase

import "todolist/m/internal/repository"

type TaskUsecase interface {
}

type taskUsecase struct {
	repo repository.TaskRepository
}

func NewTaskUsecase(repo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{repo: repo}
}
