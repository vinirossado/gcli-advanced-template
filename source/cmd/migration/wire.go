package main

import (
	"basic/pkg/logger"
	"basic/source/migration"
	"basic/source/repository"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var MigrateSet = wire.NewSet(
	migration.NewMigrate,
)

func newApp(dbType repository.DBType, conf *viper.Viper, db *gorm.DB, logger *logger.Logger) (*migration.Migrate, func(), error) {
	panic(wire.Build(
		MigrateSet,
	))
}
