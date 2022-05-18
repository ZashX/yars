package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewUnit(db *gorm.DB, namer Namer, unitName string, entity interface{}) *Unit {
	return &Unit{
		db:       db,
		namer:    namer,
		unitName: unitName,
		entity:   entity,
	}
}

type Unit struct {
	db       *gorm.DB
	namer    Namer
	unitName string
	entity   interface{}
}

func (u *Unit) Migrate(ctx context.Context) error {
	if err := u.table().AutoMigrate(u.entity); err != nil {
		return fmt.Errorf("can't migrate %s: %w", u.tableName(), err)
	}
	return nil
}

func (u *Unit) Create(ctx context.Context, data interface{}) error {
	if err := u.table().CreateInBatches(data, 1000).Error; err != nil {
		return fmt.Errorf("can't create unit %s: %w", u.tableName(), err)
	}
	return nil
}

func (u *Unit) Upsert(ctx context.Context, data interface{}) error {
	if err := u.table().Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(data, 1000).Error; err != nil {
		return fmt.Errorf("can't upsert unit %s: %w", u.tableName(), err)
	}
	return nil
}

func (u *Unit) GetUnitById(ctx context.Context, unit interface{}, id string) error {
	return u.table().Take(&unit, id).Error
}

func (u *Unit) tableName() string {
	return u.namer.unitsTable(u.unitName)
}

func (u *Unit) table() *gorm.DB {
	return u.db.Table(u.tableName())
}
