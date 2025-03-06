package repository

import (
	"context"
	"todo-app-api/task/entities"

	"gorm.io/gorm"
)

type (
	TaskRepository interface {
		Create(ctx context.Context, task *entities.Task) error
		GetAllTasks(ctx context.Context, tasks *[]entities.Task) error
	}

	TaskRepositoryImpl struct {
		db *gorm.DB
	}
)

func MakeTaskRepository(db *gorm.DB) TaskRepository {
	return &TaskRepositoryImpl{db: db}
}

func (repository TaskRepositoryImpl) Create(ctx context.Context, task *entities.Task) error {
	if err := repository.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (repository TaskRepositoryImpl) GetAllTasks(ctx context.Context, tasks *[]entities.Task) error {
	if err := repository.db.Find(tasks).Error; err != nil {
		return err
	}
	return nil
}
