package repository

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.com/yars-ai/yars/services/recommend"
	"gitlab.com/yars-ai/yars/services/recommend/mf"
	"gorm.io/gorm"
)

func NewActor(db *gorm.DB, namer Namer, recommendName string) *Actor {
	return &Actor{
		db:            db,
		namer:         namer,
		recommendName: recommendName,
	}
}

type Actor struct {
	db            *gorm.DB
	namer         Namer
	recommendName string
}

func (a *Actor) Migrate(ctx context.Context) error {
	if err := a.table().AutoMigrate(recommend.Actor{}); err != nil {
		return err
	}

	if err := a.table().AutoMigrate(mf.Actor{}); err != nil {
		return err
	}
	return nil
}

func (a *Actor) CreateActor(ctx context.Context, actor recommend.Actor) error {
	return a.table().Create(&actor).Error
}

func (a *Actor) CreateActorMF(ctx context.Context, actor mf.Actor) error {
	return a.table().Create(&actor).Error
}

func (a *Actor) UpsertActor(ctx context.Context, actor recommend.Actor) error {
	return a.table().Save(&actor).Error
}

func (a *Actor) UpsertActorMF(ctx context.Context, actor mf.Actor) error {
	return a.table().Save(&actor).Error
}

func (a *Actor) GetActorByID(ctx context.Context, id string) (recommend.Actor, error) {
	panic("not impemented")
}

func (a *Actor) GetActorMFByID(ctx context.Context, id string) (mf.Actor, error) {
	actor := mf.Actor{}
	if err := a.table().Take(&actor, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return mf.Actor{}, mf.ErrRecordNotFound
		} else {
			return mf.Actor{}, err
		}
	}
	return actor, nil
}

func (a *Actor) tableName() string {
	return a.namer.actorsTable(a.recommendName)
}

func (a *Actor) table() *gorm.DB {
	return a.db.Table(a.tableName())
}
