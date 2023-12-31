//go:build wireinject
// +build wireinject

package main

import (
	logger "basic/pkg/logger"
	"basic/source/middleware"
	"basic/source/repository"
	"basic/source/router"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var ServerSet = wire.NewSet(routes.NewServerHTTP)

var JwtSet = wire.NewSet(middleware.NewJwt)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var RepositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRepository,
	repository.NewUserRepository,
)

func newApp(repository.DBType, *viper.Viper, *logger.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
		RepositorySet,
		ServiceSet,
		HandlerSet,
		JwtSet,
	))
}
