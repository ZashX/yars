package engine

import (
	"context"
	"fmt"

	"gitlab.com/yars-ai/yars/interfaces"
	"gitlab.com/yars-ai/yars/services/recommend"
	"gitlab.com/yars-ai/yars/services/unit"

	"gitlab.com/yars-ai/yars/services/action"
)

type JSONData []byte

type Engine interface {
	GetName() string
	MigrateAll(ctx context.Context) error
	addMigrate(m interfaces.Migration)
	UnitByName(name string) (unit.Service, error)
	addUnit(name string, u unit.Service) error
	ActionByName(name string) (action.Service, error)
	addAction(name string, a action.Service) error
	RecommendByName(name string) (recommend.Service, error)
	addRecommend(name string, a recommend.Service, rateAction string) error
}

func newEngine(projectName string) Engine {
	return &engine{
		projectName: projectName,
		units:       make(map[string]unit.Service),
		actions:     make(map[string]action.Service),
		recommend:   make(map[string]recommend.Service),
	}
}

type engine struct {
	projectName string
	migrations  []interfaces.Migration
	units       map[string]unit.Service
	actions     map[string]action.Service
	recommend   map[string]recommend.Service
}

func (e *engine) GetName() string {
	return e.projectName
}

func (e *engine) MigrateAll(ctx context.Context) error {
	for _, Migrate := range e.migrations {
		if err := Migrate(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (e *engine) addMigrate(m interfaces.Migration) {
	e.migrations = append(e.migrations, m)
}

func (e *engine) UnitByName(name string) (unit.Service, error) {
	unit, ok := e.units[name]
	if !ok {
		return nil, fmt.Errorf("can't get unit by name: %s", name)
	} else {
		return unit, nil
	}
}

func (e *engine) addUnit(name string, u unit.Service) error {
	_, ok := e.units[name]
	if ok {
		return fmt.Errorf("unit already exists: %s", name)
	}
	e.units[name] = u
	return nil
}

func (e *engine) ActionByName(name string) (action.Service, error) {
	a, ok := e.actions[name]
	if !ok {
		return nil, fmt.Errorf("can't get action by name: %s", name)
	} else {
		return a, nil
	}
}

func (e *engine) addAction(name string, a action.Service) error {
	_, ok := e.actions[name]
	if ok {
		return fmt.Errorf("action already exists: %s", name)
	}
	e.actions[name] = a
	return nil
}

func (e *engine) RecommendByName(name string) (recommend.Service, error) {
	a, ok := e.recommend[name]
	if !ok {
		return nil, fmt.Errorf("can't get recommed by name: %s", name)
	} else {
		return a, nil
	}
}

func (e *engine) addRecommend(name string, r recommend.Service, actionName string) error {
	_, ok := e.recommend[name]
	if ok {
		return fmt.Errorf("recommend already exists: %s", name)
	}
	e.recommend[name] = r

	action, ok := e.actions[actionName]
	if !ok {
		return fmt.Errorf("bad action name")
	}
	action.AddHook(r.ApplyActions)

	return nil
}
