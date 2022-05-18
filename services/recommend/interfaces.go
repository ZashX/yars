package recommend

import (
	"context"
)

type CanditateGenerator interface {
	GetCanditates(ctx context.Context, actorID string, count int) ([]string, error)
}

type FeatureGenerator interface {
	GetFeatures(ctx context.Context, actionPairs []ActionPair) ([]Feature, error)
}

type Updater interface {
	FitActions(ctx context.Context, actions []Action) error
}

type Retrainer interface {
	Retrain(ctx context.Context) error
}

type RateCalculator interface {
	RateCalculate(ctx context.Context, actionPairs []ActionPair) ([]float32, error)
}
