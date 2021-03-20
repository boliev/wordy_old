// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"wordy/src/controller"
	"wordy/src/postgre"
	"wordy/src/user"
)

// Injectors from wire.go:

// InitializeUserController for controller.CreateUserController
func InitializeUserController() (controller.User, error) {
	repository, err := InitializeUserRepository()
	if err != nil {
		return controller.User{}, err
	}
	user := controller.CreateUserController(repository)
	return user, nil
}

// InitializeUserRepository wire for postgre.NewUserRepository
func InitializeUserRepository() (user.Repository, error) {
	db, err := InitializeDBConnection()
	if err != nil {
		return nil, err
	}
	repository := postgre.NewUserRepository(db)
	return repository, nil
}

// InitializeDBConnection wire for postgre.NewDBConnection
func InitializeDBConnection() (*gorm.DB, error) {
	string2 := ParameterDbConnection()
	db, err := postgre.NewDBConnection(string2)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// wire.go:

// Parameters
// InitializeDBDsn wire for db dsn
func ParameterDbConnection() string {
	return viper.GetString("database_host")
}
