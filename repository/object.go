package repository

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.com/yars-ai/yars/services/recommend"
	"gitlab.com/yars-ai/yars/services/recommend/mf"
	"gorm.io/gorm"
)

func NewObject(db *gorm.DB, namer Namer, recommendName string) *Object {
	return &Object{
		db:            db,
		namer:         namer,
		recommendName: recommendName,
	}
}

type Object struct {
	db            *gorm.DB
	namer         Namer
	recommendName string
}

func (o *Object) Migrate(ctx context.Context) error {
	if err := o.table().AutoMigrate(recommend.Object{}); err != nil {
		return err
	}
	if err := o.table().AutoMigrate(mf.Object{}); err != nil {
		return err
	}
	return nil
}

func (o *Object) CreateObject(ctx context.Context, object recommend.Object) error {
	return o.table().Create(&object).Error
}

func (o *Object) CreateObjectMF(ctx context.Context, object mf.Object) error {
	return o.table().Create(&object).Error
}

func (o *Object) UpsertObject(ctx context.Context, object recommend.Object) error {
	return o.table().Save(&object).Error
}

func (o *Object) UpsertObjectMF(ctx context.Context, object mf.Object) error {
	return o.table().Save(&object).Error
}

func (o *Object) GetObjectByID(ctx context.Context, id string) (recommend.Object, error) {
	panic("not implemented")
}

func (o *Object) GetObjectMFByID(ctx context.Context, id string) (mf.Object, error) {
	object := mf.Object{}
	if err := o.table().Take(&object, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return mf.Object{}, mf.ErrRecordNotFound
		} else {
			return mf.Object{}, err
		}
	}
	return object, nil
}

func (o *Object) GetObjectsMF(ctx context.Context) ([]mf.Object, error) {
	objects := []mf.Object{}
	err := o.table().Scan(&objects).Error
	if err != nil {
		return nil, errors.Wrap(err, "can't scan all objects mf")
	}
	return objects, nil
}

func (o *Object) tableName() string {
	return o.namer.objectsTable(o.recommendName)
}

func (o *Object) table() *gorm.DB {
	return o.db.Table(o.tableName())
}
