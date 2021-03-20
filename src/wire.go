// +build wireinject

package main

import (
	"wordy/src/controller"
	"wordy/src/postgre"
	"wordy/src/user"

	"github.com/google/wire"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// Parameters
// ParameterDbConnection wire for db dsn
func ParameterDbConnection() string {
	return viper.GetString("database_host")
}

// controllers

// InitializeUserController for controller.CreateUserController
func InitializeUserController() (controller.User, error) {
	wire.Build(controller.CreateUserController, InitializeUserRepository)
	return controller.User{}, nil
}

// InitializeUserRepository wire for postgre.NewUserRepository
func InitializeUserRepository() (user.Repository, error) {
	wire.Build(postgre.NewUserRepository, InitializeDBConnection)
	return postgre.UserRepository{}, nil
}

// InitializeDBConnection wire for postgre.NewDBConnection
func InitializeDBConnection() (*gorm.DB, error) {
	wire.Build(postgre.NewDBConnection, ParameterDbConnection)
	return &gorm.DB{}, nil
}
