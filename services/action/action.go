package action

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"

	"gitlab.com/yars-ai/yars/services/recommend"
)

type Hook func(ctx context.Context, actions []recommend.Action) error

type Service interface {
	Upsert(ctx context.Context, actions []Action) (io.Reader, error)
	AddHook(h Hook)
}

type service struct {
	name  string
	repo  Store
	hooks []Hook
}

func NewAction(name string, r Store) Service {
	return &service{
		name: name,
		repo: r,
	}
}

func (u *service) Upsert(ctx context.Context, actions []Action) (io.Reader, error) {

	if err := u.repo.Upsert(ctx, actions); err != nil {
		return nil, fmt.Errorf("can't create action %s: %w", u.name, err)
	}

	squashActions := make([]recommend.Action, len(actions))
	for i := range actions {
		squashActions[i] = recommend.Action{
			ActorID:  actions[i].ActorID,
			ObjectID: actions[i].ObjectID,
			Rate:     actions[i].Rate,
		}
	}

	for _, v := range u.hooks {
		if err := v(ctx, squashActions); err != nil {
			log.Printf("can't hook: %v", err)
		}
	}

	return bytes.NewReader([]byte{}), nil
}

func (u *service) AddHook(h Hook) {
	u.hooks = append(u.hooks, h)
}

type Action struct {
	ActorID  string  `json:"actor_id" gorm:"size:256;primaryKey"`
	ObjectID string  `json:"object_id" gorm:"size:256;primaryKey"`
	Rate     float32 `json:"rate"`
}
