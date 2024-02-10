// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"basic/pkg/logger"
	"basic/source/handler"
	"basic/source/middleware"
	"basic/source/repository"
	"basic/source/router"
	"basic/source/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Injectors from wire.go:

func newApp(dbType repository.DBType, viperViper *viper.Viper, loggerLogger *logger.Logger) (*gin.Engine, func(), error) {
	jwt := middleware.NewJwt(viperViper)
	handlerHandler := handler.NewHandler(loggerLogger)
	serviceService := service.NewService(loggerLogger, jwt)
	db := repository.NewDB(dbType, viperViper)
	repositoryRepository := repository.NewRepository(loggerLogger, db)
	userRepository := repository.NewUserRepository(repositoryRepository)
	userService := service.NewUserService(serviceService, userRepository)
	userHandler := handler.NewUserHandler(handlerHandler, userService)
	engine := routes.NewServerHTTP(loggerLogger, jwt, userHandler)
	return engine, func() {
	}, nil
}

// wire.go:

var ServerSet = wire.NewSet(routes.NewServerHTTP)

var JwtSet = wire.NewSet(middleware.NewJwt)

var RepositorySet = wire.NewSet(repository.NewDB, repository.NewRepository, repository.NewUserRepository)

var ServiceSet = wire.NewSet(service.NewService, service.NewUserService)

var HandlerSet = wire.NewSet(handler.NewHandler, handler.NewUserHandler)