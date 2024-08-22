//go:build wireinject
// +build wireinject

package main

import (
	"basic/pkg/app"
	"basic/pkg/logger"
	"basic/source/migration"
	"basic/source/repository"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	repository.NewRepository,
	repository.NewUserRepository,
)
var serverSet = wire.NewSet(
	server.NewMigrate,
)

// build App
func newApp(
	migrate *server.Migrate,
) *app.App {
	return app.NewApp(
		app.WithServer(migrate),
		app.WithName("demo-migrate"),
	)
}

func NewWire(*viper.Viper, *logger.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serverSet,
		newApp,
	))
}
