package main

import (
	"wordy/src/postgre"
	"wordy/src/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ProvideUserRepository implementation coustructor for UserRepository by  go.uber.org/fx
func ProvideUserRepository() user.Repository {
	rep := postgre.UserRepository{}
	return rep
}

// ProvideDbService returns DB connection
func ProvideDbService() *gorm.DB {
	dsn := "host=localhost user=wordy password=123456 dbname=wordy port=5433 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to the database")
	}

	return db
}
