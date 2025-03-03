package server

import (
	"fmt"
	"todo-app-api/config"
	"todo-app-api/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type EchoServer struct {
	database *database.PostgresDatabase
	config   *config.Config
	app      *echo.Echo
}

func (server *EchoServer) Start() {
	server.app.Use(middleware.Recover())
	server.app.Use(middleware.Logger())

	server.app.GET("healt", func(c echo.Context) error {
		return c.String(200, "Welcome to healt service")
	})

	serverUrl := fmt.Sprintf(":%d", server.config.Server.Port)
	server.app.Logger.Fatal(server.app.Start(serverUrl))
}

func MakeNewEchoServer(config *config.Config, database *database.PostgresDatabase) Server {
	newEchoApp := echo.New()
	newEchoApp.Logger.SetLevel(log.DEBUG)

	return &EchoServer{
		database: database,
		config:   config,
		app:      newEchoApp,
	}
}
