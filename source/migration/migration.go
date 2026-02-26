package server

import (
	"context"
	"fmt"
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
	models := model.RetrieveAll()

	if err := m.db.AutoMigrate(models...); err != nil {
		m.log.Error("user migrate error", zap.Error(err))
		return err
	}

	for _, mdl := range models {
		if err := m.syncRemovedColumns(mdl); err != nil {
			m.log.Error("syncRemovedColumns error", zap.Error(err))
			return err
		}
	}

	m.log.Info("AutoMigrate success")
	os.Exit(0)
	return nil
}

// syncRemovedColumns drops DB columns that no longer exist in the struct.
// GORM's AutoMigrate only adds/modifies columns — it never removes them.
func (m *Migrate) syncRemovedColumns(mdl interface{}) error {
	if !m.db.Migrator().HasTable(mdl) {
		return nil
	}

	stmt := &gorm.Statement{DB: m.db}
	if err := stmt.Parse(mdl); err != nil {
		return fmt.Errorf("parse model: %w", err)
	}

	columnTypes, err := m.db.Migrator().ColumnTypes(mdl)
	if err != nil {
		return fmt.Errorf("get column types for %s: %w", stmt.Table, err)
	}

	structCols := make(map[string]bool)
	for _, field := range stmt.Schema.Fields {
		if field.DBName != "" {
			structCols[field.DBName] = true
		}
	}

	for _, col := range columnTypes {
		if !structCols[col.Name()] {
			m.log.Info("Dropping removed column",
				zap.String("table", stmt.Table),
				zap.String("column", col.Name()),
			)
			if err := m.db.Migrator().DropColumn(mdl, col.Name()); err != nil {
				return fmt.Errorf("drop column %s.%s: %w", stmt.Table, col.Name(), err)
			}
		}
	}

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
