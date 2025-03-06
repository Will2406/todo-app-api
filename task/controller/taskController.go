package controller

import (
	"net/http"
	"time"
	"todo-app-api/core"
	"todo-app-api/task/entities"
	"todo-app-api/task/usecase"

	"github.com/labstack/echo/v4"
)

type (
	TaskController interface {
		GetAllTasks(ctx echo.Context) error
		Create(ctx echo.Context) error
	}

	TaskControllerImpl struct {
		usecase usecase.TaskUseCase
	}
)

func MakeTaskController(usecase usecase.TaskUseCase) TaskController {
	return &TaskControllerImpl{usecase: usecase}
}

func (controller TaskControllerImpl) Create(ctx echo.Context) error {
	task := entities.TaskBase{
		UpdatedAt: time.Now(),
	}

	if err := ctx.Bind(&task); err != nil {
		return ctx.JSON(http.StatusBadRequest, core.HandleError(err))
	}

	if err := controller.usecase.Create(ctx.Request().Context(), task); err != nil {
		return ctx.JSON(http.StatusInternalServerError, core.HandleError(err))
	}

	return ctx.JSON(http.StatusCreated, core.BaseResponse{
		Message: "Success",
		Status:  "Success",
	})

}

func (controller TaskControllerImpl) GetAllTasks(ctx echo.Context) error {
	taskResponse, err := controller.usecase.GetAllTasks(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, core.HandleError(err))
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
