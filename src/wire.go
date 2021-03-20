// +build wireinject

package main

import (
	"wordy/src/controller"
	"wordy/src/postgre"
	"wordy/src/user"

	"github.com/google/wire"
	"gorm.io/gorm"
)

// Parameters

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
	wire.Build(postgre.NewDBConnection, "host=localhost user=wordy password=123456 dbname=wordy port=5433 sslmode=disable TimeZone=Europe/Moscow")
	return &gorm.DB{}, nil
}
