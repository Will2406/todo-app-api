package usecase

import (
	"context"
	"todo-app-api/task/entities"
	"todo-app-api/task/repository"
)

type (
	TaskUseCase interface {
		Create(ctx context.Context, task entities.TaskBase) error
		GetAllTasks(ctx context.Context) ([]entities.TaskBase, error)
		Update(ctx context.Context, task entities.TaskBase) error
		Delete(ctx context.Context, id uint) error
	}

	TaskUseCaseImpl struct {
		repository repository.TaskRepository
	}
)

func MakeTaskUseCase(rp repository.TaskRepository) TaskUseCase {
	return &TaskUseCaseImpl{repository: rp}
}

func (usecase TaskUseCaseImpl) Create(ctx context.Context, task entities.TaskBase) error {
	return usecase.repository.Create(ctx, task.MapToModel())
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
func (usecase *TaskUseCaseImpl) Update(ctx context.Context, task entities.TaskBase) error {
	taskModel := task.MapToModel()
	return usecase.repository.Update(ctx, taskModel)
}

func (usecase *TaskUseCaseImpl) Delete(ctx context.Context, id uint) error {
	return usecase.repository.Delete(ctx, id)
}
