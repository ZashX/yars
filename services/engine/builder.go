package engine

import (
	"context"
	"fmt"

	"gitlab.com/yars-ai/yars/ast"
	"gitlab.com/yars-ai/yars/interfaces"
	"gitlab.com/yars-ai/yars/services/action"
	"gitlab.com/yars-ai/yars/services/entity"
	"gitlab.com/yars-ai/yars/services/recommend"
	"gitlab.com/yars-ai/yars/services/recommend/boosting"
	"gitlab.com/yars-ai/yars/services/recommend/feature"
	"gitlab.com/yars-ai/yars/services/recommend/mf"
	"gitlab.com/yars-ai/yars/services/unit"
)

type engineBuilder struct {
	repo    Fabric
	project *ast.Project
}

func NewBuilder(repo Fabric, project *ast.Project) *engineBuilder {
	return &engineBuilder{
		repo:    repo,
		project: project,
	}
}

func (b *engineBuilder) Build() (Engine, error) {
	e := newEngine(b.project.Name.Value)

	for _, s := range b.project.Statements {
		switch t := s.(type) {
		case *ast.UnitStatement:
			u, name, mg, err := b.buildUnit(t)
			if err != nil {
				return nil, fmt.Errorf("can't build unit: %w", err)
			}
			if err = e.addUnit(name, u); err != nil {
				return nil, fmt.Errorf("can't add unit: %w", err)
			}
			e.addMigrate(mg)
		case *ast.ActionStatement:
			a, name, mg, err := b.buildAction(t)
			if err != nil {
				return nil, fmt.Errorf("can't build action: %w", err)
			}
			if err = e.addAction(name, a); err != nil {
				return nil, fmt.Errorf("can't add action: %w", err)
			}
			e.addMigrate(mg)
		case *ast.RecommendStatement:
			continue
		default:
			return nil, fmt.Errorf("unexpected statement type: %T", t)
		}
	}

	for _, s := range b.project.Statements {
		switch t := s.(type) {
		case *ast.UnitStatement:
			continue
		case *ast.ActionStatement:
			continue
		case *ast.RecommendStatement:
			b, err := b.buildRecommend(t, e)
			if err != nil {
				return nil, fmt.Errorf("can't build recommend service: %w", err)
			}
			if err = e.addRecommend(b.name, b.service, b.action); err != nil {
				return nil, fmt.Errorf("can't add recommend service: %w", err)
			}
			e.addMigrate(b.migration)
		default:
			return nil, fmt.Errorf("unexpected statement type: %T", t)
		}
	}
	return e, nil
}

func (b engineBuilder) buildAction(a *ast.ActionStatement) (action.Service, string, interfaces.Migration, error) {
	aName := a.Name.Value
	repo, migration := b.repo.Action(aName)
	aController := action.NewAction(aName, repo)
	return aController, aName, migration, nil
}

func (b engineBuilder) buildUnit(u *ast.UnitStatement) (unit.Service, string, interfaces.Migration, error) {

	eBuilder := entity.NewEntityBuilder()

	for _, field := range u.FieldsDeclaration.Fields {
		tag := fmt.Sprintf("json:\"%s\"", field.Name.Value)

		if err := eBuilder.AddField(field.Name.Value, field.Type, tag); err != nil {
			return nil, "", nil, fmt.Errorf("can't add field %s: %w", field.Name.Value, err)
		}
	}

	eConstructor := eBuilder.BuildUnit()
	uName := u.Name.Value
	unitRepo, migration := b.repo.Unit(uName, eConstructor.New())
	uController := unit.NewUnit(uName, unitRepo, eConstructor)

	return uController, uName, migration, nil
}

func (b engineBuilder) buildRecommend(u *ast.RecommendStatement, e Engine) (*buildRecommendRes, error) {
	rName := u.Name.Value

	actorMFStore, migrationActorMf := b.repo.ActorMF(rName)
	objectMFStore, migrationObjectMf := b.repo.ObjectMF(rName)
	squashActionStore, migrationSquashMF := b.repo.SquashAction(rName)

	als := mf.NewALS("als", squashActionStore, actorMFStore, objectMFStore)

	actorUnit, err := e.UnitByName(u.Actor.Value)
	if err != nil {
		return nil, err
	}
	objectUnit, err := e.UnitByName(u.Object.Value)
	if err != nil {
		return nil, err
	}
	unitFeatureGenerator := feature.NewUnitFeatureGenerator(actorUnit, objectUnit)
	mergedFeatureGenerator := feature.NewMergedFeatureGenerator([]recommend.FeatureGenerator{unitFeatureGenerator})

	cb := boosting.NewCatboost("catboost", "../train", squashActionStore, mergedFeatureGenerator, als)

	rec := recommend.NewService(
		"rec_service",
		[]recommend.Updater{als},
		[]recommend.Retrainer{als, cb},
		cb,
		squashActionStore,
	)

	migration := func(ctx context.Context) error {
		if err := migrationActorMf(ctx); err != nil {
			return err
		}
		if err := migrationObjectMf(ctx); err != nil {
			return err
		}
		if err := migrationSquashMF(ctx); err != nil {
			return err
		}
		return nil
	}

	action, err := findRateAction(*u.ParamsDeclaration)
	if err != nil {
		return nil, err
	}

	return &buildRecommendRes{
		name:      rName,
		service:   rec,
		migration: migration,
		action:    action,
	}, nil
}

type buildRecommendRes struct {
	name      string
	service   recommend.Service
	migration interfaces.Migration
	action    string
}

//TODO MADE ACTIONS
func findRateAction(p ast.ParamsDeclaration) (string, error) {
	for _, v := range p.Params {
		if v.Name.Value != "rate" {
			continue
		}
		if ident, ok := v.Expression.(*ast.Lvalue); ok {
			return ident.Name.Value, nil
		}
		return "", fmt.Errorf("bad rate value, expected one ast.Identifier, got %+v", v.Expression)
	}

	return "", fmt.Errorf("not exist required field rate")
}
