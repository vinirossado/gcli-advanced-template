//go:build wireinject
// +build wireinject

package main

import (
	"log"
	"net/http"

	"github.com/google/wire"
	"github.com/spf13/viper"

	"basic/source/handler"
	"basic/source/middleware"
	"basic/source/repository"
	routes "basic/source/router"
	"basic/source/service"
)

var serverSet = wire.NewSet(routes.NewServerHTTP)

var JwtSet = wire.NewSet(middleware.NewJwt)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
)

func newApp(
	httpServer *http.Server,
) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		// sid.NewSid,
		// jwt.NewJwt,
		newApp,
	))
}
