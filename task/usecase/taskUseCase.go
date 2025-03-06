package usecase

import (
	"context"
	"todo-app-api/task/entities"
	"todo-app-api/task/repository"
)

type (
	TaskUseCase interface {
		Create() error
		GetAllTasks(ctx context.Context) ([]entities.TaskBase, error)
	}

	TaskUseCaseImpl struct {
		repository repository.TaskRepository
	}
)

func MakeTaskUseCase(rp repository.TaskRepository) TaskUseCase {
	return &TaskUseCaseImpl{repository: rp}
}

func (usecase TaskUseCaseImpl) Create() error {
	return nil
}

func (usecase TaskUseCaseImpl) GetAllTasks(ctx context.Context) ([]entities.TaskBase, error) {
	tasks := []entities.Task{}
	if err := usecase.repository.GetAllTasks(ctx, &tasks); err != nil {
		return nil, err
	}

	taskResp := []entities.TaskBase{}
	for _, task := range tasks {
		taskResp = append(taskResp, task.MapToBase())
	}
	return taskResp, nil
}
