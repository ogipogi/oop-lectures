// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package config

import (
	"example-rest-api/app/controller"
	"example-rest-api/app/repository"
	"example-rest-api/app/service"
	"github.com/google/wire"
)

// Injectors from injector.go:

func Init() *Initialization {
	gormDB := ConnectToDB()
	userRepositoryImpl := repository.UserRepositoryInit(gormDB)
	userServiceImpl := service.UserServiceInit(userRepositoryImpl)
	userControllerImpl := controller.UserControllerInit(userServiceImpl)
	initialization := NewInitialization(userRepositoryImpl, userServiceImpl, userControllerImpl)
	return initialization
}

// injector.go:

var db = wire.NewSet(ConnectToDB)

var userServiceSet = wire.NewSet(service.UserServiceInit, wire.Bind(new(service.UserService), new(*service.UserServiceImpl)))

var userRepoSet = wire.NewSet(repository.UserRepositoryInit, wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)))

var userCtrlSet = wire.NewSet(controller.UserControllerInit, wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)))
