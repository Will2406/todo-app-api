package main

import (
	"todo-app-api/config"
	"todo-app-api/database"
	"todo-app-api/task/entities"
)

func main() {
	conf := config.GetConfig()
	db := database.NewConnectionToPostgressDatabase(conf)

	makeMigrations(db)
}

func makeMigrations(database *database.PostgresDatabase) {
	database.Db.AutoMigrate(&entities.Task{})
	database.Db.CreateInBatches([]entities.Task{
		{
			Name: "Work",
		},
		{
			Name: "Gym",
		},
		{
			Name: "Make breakfast",
		},
		{
			Name: "Make dinner",
		},
	}, 5)
}
