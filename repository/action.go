package repository

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/yars-ai/yars/services/action"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewAction(db *gorm.DB, namer Namer, actionName string) *Action {
	return &Action{
		db:         db,
		namer:      namer,
		actionName: actionName,
	}
}

type Action struct {
	db         *gorm.DB
	namer      Namer
	actionName string
}

func (u *Action) Migrate(ctx context.Context) error {
	if err := u.table().AutoMigrate(actionDTO{}); err != nil {
		return fmt.Errorf("can't migrate %s: %w", u.tableName(), err)
	}
	return nil
}

func (u *Action) Upsert(ctx context.Context, actions []action.Action) error {
	var data []actionDTO
	for _, v := range actions {
		data = append(data, actionDTO{
			ActorID:  v.ActorID,
			ObjectID: v.ObjectID,
			Name:     u.actionName,
			Rate:     v.Rate,
		})
	}

	if err := u.table().Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(data, 1000).Error; err != nil {
		return fmt.Errorf("can't create unit %s: %w", u.tableName(), err)
	}
	return nil
}

func (u *Action) tableName() string {
	return u.namer.actionsTable()
}

func (u *Action) table() *gorm.DB {
	return u.db.Table(u.tableName())
}

type actionDTO struct {
	ActorID   string    `db:"actor_id" gorm:"size:256;primaryKey"`
	ObjectID  string    `db:"object_id" gorm:"size:256;primaryKey"`
	Name      string    `db:"name" gorm:"size:256;primaryKey"`
	Rate      float32   `db:"rate"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"created_at"`
}
