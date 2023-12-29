package repository

import (
	"basic/pkg/logger"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type DBType string

const (
	SqlServer  DBType = "data.sqlserver.connectionString"
	PostgreSQL DBType = "data.postgresql.connectionString"
)

type Repository struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewRepository(logger *logger.Logger, db *gorm.DB) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

func NewDB(dbType DBType, conf *viper.Viper) *gorm.DB {
	var db *gorm.DB
	if dbType == PostgreSQL {
		db, _ = connectPostgresql(conf)
	}
	db, _ = connectSqlServer(conf)

	return db
}

func connectSqlServer(conf *viper.Viper) (*gorm.DB, error) {
	db, err := gorm.Open(sqlserver.Open(conf.GetString(string(SqlServer))), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db, nil
}

func connectPostgresql(conf *viper.Viper) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(conf.GetString(string(PostgreSQL))), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open PostgreSQL database: %w", err)
	}
	defer db.Close()
	return db, nil
}
