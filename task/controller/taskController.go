package controller

import (
	"net/http"
	"strconv"
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
		Update(ctx echo.Context) error
		Delete(ctx echo.Context) error
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

func (controller TaskControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	taskID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, core.HandleError(err))
	}

	task := entities.TaskBase{
		ID:        parseUint(id),
		UpdatedAt: time.Now(),
	}

	if err := ctx.Bind(&task); err != nil {
		return ctx.JSON(http.StatusBadRequest, core.HandleError(err))
	}

	task.ID = uint(taskID)

	if err := controller.usecase.Update(ctx.Request().Context(), task); err != nil {
		return ctx.JSON(http.StatusInternalServerError, core.HandleError(err))
	}

	return ctx.JSON(http.StatusOK, core.BaseResponse{
		Message: "Task updated successfully",
		Status:  "Success",
	})
}

func (controller TaskControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")

	if err := controller.usecase.Delete(ctx.Request().Context(), parseUint(id)); err != nil {
		return ctx.JSON(http.StatusBadRequest, core.HandleError(err))
	}

	return ctx.JSON(http.StatusOK, core.BaseResponse{
		Message: "Task deleted successfully",
		Status:  "Success",
	})
}

func parseUint(s string) uint {
	id, _ := strconv.ParseUint(s, 10, 32)
	return uint(id)
}
