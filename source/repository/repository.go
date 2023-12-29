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
	SqlServer  DBType = "sqlserver"
	PostgreSQL DBType = "postgresql"
)

type Database interface {
	Connect() (*gorm.DB, error)
}

type SQLServerDatabase struct {
	connectionString string
}

type PostgreSQLDatabase struct {
	connectionString string
}

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

func NewDB(dbType DBType, conf *viper.Viper) (Database, error) {
	switch dbType {
	case SqlServer:
		return &SQLServerDatabase{
			connectionString: conf.GetString("data.sqlserver.connectionString"),
		}, nil
	case PostgreSQL:
		return &PostgreSQLDatabase{
			connectionString: conf.GetString("data.postgresql.connectionString"),
		}, nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}

func (s *SQLServerDatabase) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlserver.Open(s.connectionString), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})

	if err != nil {
		return nil, err
	}
	
	return db, nil
}

func (p *PostgreSQLDatabase) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(p.connectionString), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
