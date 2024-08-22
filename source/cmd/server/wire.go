//go:build wireinject
// +build wireinject

package main

import (
	"basic/pkg/app"
	"basic/pkg/logger"
	"basic/pkg/server/http"

	"github.com/google/wire"
	"github.com/spf13/viper"

	"basic/pkg/jwt"
	"basic/source/handler"
	"basic/source/repository"
	routes "basic/source/router"
	"basic/source/service"
)

var serverSet = wire.NewSet(routes.NewHTTPServer)

var JwtSet = wire.NewSet(jwt.NewJwt)

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

func newApp(httpServer *http.Server) *app.App {
	return app.NewApp(
		app.WithServer(httpServer),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *logger.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		// sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
