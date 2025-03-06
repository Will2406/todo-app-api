package entities

import (
	"time"
	"todo-app-api/core"
)

type Task struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"not null"`
	CreateAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TaskBase struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TasksResponse struct {
	core.BaseResponse
	Tasks []TaskBase `json:"data"`
}

func (t *Task) MapToBase() TaskBase {
	return TaskBase{
		ID:        t.ID,
		Name:      t.Name,
		UpdatedAt: t.UpdatedAt,
	}
}

func (t *TaskBase) MapToModel() *Task {
	return &Task{
		ID:        t.ID,
		Name:      t.Name,
		UpdatedAt: t.UpdatedAt,
	}
}
