package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"

	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"basic/pkg/logger"
	"basic/pkg/zapgorm2"
)

type DBType string

const ctxTxKey = "TxKey"

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

type Transaction interface {
	Transaction(ctx context.Context, fn func(ctx context.Context) error) error
}

func NewTransaction(r *Repository) Transaction {
	return r
}

func (r *Repository) DB(ctx context.Context) *gorm.DB {
	v := ctx.Value(ctxTxKey)
	if v != nil {
		if tx, ok := v.(*gorm.DB); ok {
			return tx
		}
	}
	return r.db.WithContext(ctx)
}

func (r *Repository) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, ctxTxKey, tx)
		return fn(ctx)
	})
}

func NewDB(conf *viper.Viper, l *logger.Logger) *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)

	log := zapgorm2.New(l.Logger)
	driver := conf.GetString("data.db.user.driver")
	dsn := conf.GetString("data.db.user.dsn")

	// GORM doc: https://gorm.io/docs/connecting_to_the_database.html
	switch driver {
	case "sqlserver":
		db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
			Logger: log,
		})
	case "postgres":
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	default:
		panic("unknown db driver")
	}
	if err != nil {
		panic(err)
	}
	db = db.Debug()

	// Connection Pool config
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}

func connectSqlServer(conf *viper.Viper) (*gorm.DB, error) {
	fmt.Println(conf.GetString(string(SqlServer)))

	db, err := gorm.Open(sqlserver.Open(conf.GetString(string(SqlServer))), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		panic(err)
	}

	return db, nil
}

func connectPostgresql(conf *viper.Viper) (*gorm.DB, error) {
	fmt.Println(conf.GetString(string(PostgreSQL)))
	db, err := gorm.Open(postgres.Open(conf.GetString(string(PostgreSQL))), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open PostgreSQL database: %w", err)
	}
	return db, nil
}
