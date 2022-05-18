package repository

import (
	"gitlab.com/yars-ai/yars/interfaces"
	"gitlab.com/yars-ai/yars/services/action"
	"gitlab.com/yars-ai/yars/services/recommend"
	"gitlab.com/yars-ai/yars/services/recommend/mf"
	"gitlab.com/yars-ai/yars/services/unit"
	"gorm.io/gorm"
)

type fabric struct {
	db    *gorm.DB
	namer Namer
}

func NewFabric(db *gorm.DB, projectName string) *fabric {
	return &fabric{db: db, namer: NewNamer(projectName)}
}

func (f *fabric) Unit(unitName string, entity interface{}) (unit.Store, interfaces.Migration) {
	u := NewUnit(f.db, f.namer, unitName, entity)
	return u, u.Migrate
}

func (f *fabric) Action(actionName string) (action.Store, interfaces.Migration) {
	a := NewAction(f.db, f.namer, actionName)
	return a, a.Migrate
}

func (f *fabric) Actor(recommendName string) (recommend.ActorStore, interfaces.Migration) {
	a := NewActor(f.db, f.namer, recommendName)
	return a, a.Migrate
}

func (f *fabric) ActorMF(recommendName string) (mf.ActorStore, interfaces.Migration) {
	a := NewActor(f.db, f.namer, recommendName)
	return a, a.Migrate
}

func (f *fabric) Object(recommendName string) (recommend.ObjectStore, interfaces.Migration) {
	o := NewObject(f.db, f.namer, recommendName)
	return o, o.Migrate
}

func (f *fabric) ObjectMF(recommendName string) (mf.ObjectStore, interfaces.Migration) {
	o := NewObject(f.db, f.namer, recommendName)
	return o, o.Migrate
}

func (f *fabric) SquashAction(recommendName string) (recommend.SquashActionStore, interfaces.Migration) {
	a := NewSquashAction(f.db, f.namer, recommendName)
	return a, a.Migrate
}

func NewFabricMock(projectName string) *fabric {
	return &fabric{db: nil, namer: NewNamer(projectName)}
}
