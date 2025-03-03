package database

import (
	"fmt"
	"sync"
	"todo-app-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDatabase struct {
	Db *gorm.DB
}

var (
	once             sync.Once
	databaseInstance *PostgresDatabase
)

func NewConnectionToPostgressDatabase(conf *config.Config) *PostgresDatabase {
	once.Do(
		func() {
			dsn := fmt.Sprintf(
				"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
				conf.Db.Host,
				conf.Db.User,
				conf.Db.Password,
				conf.Db.Name,
				conf.Db.Port,
				conf.Db.SSLMode,
			)

			db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

			if err != nil {
				panic("Error connection to database")
			}
			databaseInstance = &PostgresDatabase{Db: db}
		},
	)
	return databaseInstance
}

func (database *PostgresDatabase) GetDb() *gorm.DB {
	return database.Db
}
