package server

import (
	"context"
	"os"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"basic/pkg/logger"
	"basic/source/model"
)

type Migrate struct {
	db  *gorm.DB
	log *logger.Logger
}

func NewMigrate(db *gorm.DB, log *logger.Logger) *Migrate {
	return &Migrate{
		db:  db,
		log: log,
	}
}
func (m *Migrate) Start(ctx context.Context) error {
	if err := m.db.AutoMigrate(model.RetrieveAll()...); err != nil {
		m.log.Error("user migrate error", zap.Error(err))
		return err
	}
	m.log.Info("AutoMigrate success")
	os.Exit(0)
	return nil
}
func (m *Migrate) Stop(ctx context.Context) error {
	m.log.Info("AutoMigrate stop")
	return nil
}

func (m *Migrate) DropAll() {
	err := m.db.Migrator().DropTable(model.RetrieveAll()...)
	if err != nil {
		return
	}
}
