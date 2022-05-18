package recommend

import (
	"context"

	"github.com/pkg/errors"
)

type Service interface {
	GetTopRelevant(ctx context.Context, actorID string, count int) ([]string, error)
	ApplyActions(ctx context.Context, actions []Action) error
	Retrain(ctx context.Context) error
	GetRetrainers() []Retrainer
}

func NewService(
	name string,
	u []Updater,
	r []Retrainer,
	c CanditateGenerator,
	sqActions SquashActionStore,
) *service {
	rs := recommendState{
		isTrained: false,
	}
	return &service{
		name:           name,
		recommendState: rs,
		retrainers:     r,
		updaters:       u,
		candidates:     c,
		sqActions:      sqActions,
	}
}

type recommendState struct {
	isTrained bool
}

type service struct {
	name           string
	recommendState recommendState
	sqActions      SquashActionStore
	updaters       []Updater
	retrainers     []Retrainer
	candidates     CanditateGenerator
}

func (r *service) GetRetrainers() []Retrainer {
	return r.retrainers
}

func (r *service) GetTopRelevant(ctx context.Context, actorID string, count int) ([]string, error) {
	candidates, err := r.candidates.GetCanditates(ctx, actorID, count)
	if err != nil {
		return nil, err
	}
	return candidates, nil
}

func (r *service) ApplyActions(ctx context.Context, actions []Action) error {
	if err := r.sqActions.Upsert(ctx, actions); err != nil {
		return errors.Wrapf(err, "can't apply squash actions for recommend %s", r.name)
	}

	if r.recommendState.isTrained {
		for i := range r.updaters {
			err := r.updaters[i].FitActions(ctx, actions)
			if err != nil {
				return errors.Wrap(err, "can't fit actions")
			}
		}
	}
	return nil
}

func (r *service) Retrain(ctx context.Context) error {
	for _, r := range r.retrainers {
		if err := r.Retrain(ctx); err != nil {
			return err
		}
	}
	r.recommendState.isTrained = true
	return nil
}
