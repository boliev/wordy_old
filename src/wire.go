// +build wireinject

package main

import (
	"wordy/src/controller"
	"wordy/src/postgre"

	"github.com/google/wire"
	"gorm.io/gorm"
)

// controllers

// InitializeUserController for controller.CreateUserController
func InitializeUserController() (controller.User, error) {
	wire.Build(controller.CreateUserController, InitializeUserRepository)
	return controller.User{}, nil
}

// InitializeUserRepository wire for postgre.NewUserRepository
func InitializeUserRepository() (postgre.UserRepository, error) {
	wire.Build(postgre.NewUserRepository, postgre.NewDBConnection)
	return postgre.UserRepository{}, nil
}

// InitializeDBConnection wire for postgre.NewDBConnection
func InitializeDBConnection() (*gorm.DB, error) {
	wire.Build(postgre.NewDBConnection)
	return &gorm.DB{}, nil
}
