package router

import (
	"todo-app-api/task/controller"
	"todo-app-api/task/repository"
	"todo-app-api/task/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func MakeTaskRouter(database *gorm.DB, server *echo.Echo) {
	repository := repository.MakeTaskRepository(database)
	usecase := usecase.MakeTaskUseCase(repository)
	controller := controller.MakeTaskController(usecase)

	server.GET("/tasks", controller.GetAllTasks)
	server.POST("/task", controller.Create)
}
