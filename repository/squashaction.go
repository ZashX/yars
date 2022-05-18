package repository

import (
	"context"

	werrors "github.com/pkg/errors"
	"gitlab.com/yars-ai/yars/services/recommend"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewSquashAction(db *gorm.DB, namer Namer, recommendName string) *SquashAction {
	return &SquashAction{
		db:            db,
		namer:         namer,
		recommendName: recommendName,
	}
}

type SquashAction struct {
	db            *gorm.DB
	namer         Namer
	recommendName string
}

func (u *SquashAction) Migrate(ctx context.Context) error {
	if err := u.table().AutoMigrate(recommend.Action{}); err != nil {
		return werrors.Wrapf(err, "can't migrate %s: %w", u.tableName())
	}
	return nil
}

func (u *SquashAction) Upsert(ctx context.Context, actions []recommend.Action) error {
	if err := u.table().Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(actions, 1000).Error; err != nil {
		return werrors.Wrapf(err, "can't create unit %s: %w", u.tableName())
	}
	return nil
}

func (u *SquashAction) ScanAll(ctx context.Context) ([]recommend.Action, error) {
	actions := []recommend.Action{}
	err := u.table().Find(&actions).Error
	if err != nil {
		return []recommend.Action{}, werrors.Wrap(err, "can't scan all actions")
	}
	return actions, nil
}

func (u *SquashAction) tableName() string {
	return u.namer.squashActionsTable(u.recommendName)
}

func (u *SquashAction) table() *gorm.DB {
	return u.db.Table(u.tableName())
}
