//go:build wireinject
// +build wireinject

package main

import (
	"basic/pkg/logger"
	"basic/source/migration"
	"basic/source/repository"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var RepositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRepository,
	repository.NewUserRepository,
)

var MigrateSet = wire.NewSet(
	migration.NewMigrate,
)

func newApp(*viper.Viper, *logger.Logger) (*migration.Migrate, func(), error) {
	panic(wire.Build(
		RepositorySet,
		MigrateSet,
	))
}
