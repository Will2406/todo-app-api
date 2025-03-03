package main

import (
	"todo-app-api/config"
	"todo-app-api/database"
	"todo-app-api/server"
)

func main() {
	config := config.GetConfig()
	db := database.NewConnectionToPostgressDatabase(config)
	server.MakeNewEchoServer(config, db).Start()
}
