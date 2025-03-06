package controller

import (
	"net/http"
	"todo-app-api/core"
	"todo-app-api/task/entities"
	"todo-app-api/task/usecase"

	"github.com/labstack/echo/v4"
)

type (
	TaskController interface {
		GetAllTasks(ctx echo.Context) error
	}

	TaskControllerImpl struct {
		usecase usecase.TaskUseCase
	}
)

func MakeTaskController(usecase usecase.TaskUseCase) TaskController {
	return &TaskControllerImpl{usecase: usecase}
}

func (controller TaskControllerImpl) GetAllTasks(ctx echo.Context) error {
	taskResponse, err := controller.usecase.GetAllTasks(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	response := entities.TasksResponse{
		BaseResponse: core.BaseResponse{
			Message: "goo",
			Status:  "wuu",
		},
		Tasks: taskResponse,
	}
	return ctx.JSON(http.StatusOK, response)
}
