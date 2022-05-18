package mf

import (
	"context"
	"errors"

	werrors "github.com/pkg/errors"
	"gitlab.com/yars-ai/yars/services/recommend"
	"gitlab.com/yars-ai/yars/src"
)

type ALS struct {
	name        string
	actionStore recommend.SquashActionStore
	actorStore  ActorStore
	objectStore ObjectStore
	rng         src.RandomGenerator
	nFactors    int
	initMean    float32
	initStd     float32
	lr          float32
}

func NewALS(
	name string,
	actionStore recommend.SquashActionStore,
	actorStore ActorStore,
	objectStore ObjectStore,
) *ALS {
	return &ALS{
		name:        name,
		actionStore: actionStore,
		actorStore:  actorStore,
		objectStore: objectStore,
		rng:         src.NewRandomGenerator(42),
		nFactors:    10,
		initMean:    0,
		initStd:     1,
		lr:          0.03,
	}
}

func rateCalc(a Actor, o Object) float32 {
	return a.Bias + o.Bias + src.Dot(a.Vector, o.Vector)
}

func (als *ALS) gradStep(a *Actor, o *Object, rate float32) {
	diff := a.Bias + o.Bias + src.Dot(a.Vector, o.Vector) - rate
	a.Bias -= 2 * diff * als.lr
	o.Bias -= 2 * diff * als.lr
	a.Vector = src.NormIf(src.SubVV(a.Vector, src.MulVS(o.Vector, diff*als.lr)))
	o.Vector = src.NormIf(src.SubVV(o.Vector, src.MulVS(a.Vector, diff*als.lr)))
}

func (als *ALS) newActor(actorID string) Actor {
	return Actor{
		Actor:  recommend.Actor{ID: actorID},
		Vector: als.rng.NewNormalVector(als.nFactors, als.initMean, als.initStd),
		Bias:   0,
	}
}

func (als *ALS) newObject(objectID string) Object {
	return Object{
		Object: recommend.Object{ID: objectID},
		Vector: als.rng.NewNormalVector(als.nFactors, als.initMean, als.initStd),
		Bias:   0,
	}
}

func (als *ALS) GetCanditates(ctx context.Context, actorID string, count int) ([]string, error) {
	actor, err := als.actorStore.GetActorMFByID(ctx, actorID)
	if err != nil {
		return nil, err
	}
	objects, err := als.objectStore.GetObjectsMF(ctx)
	if err != nil {
		return nil, err
	}
	rates := make([]float32, len(objects))
	candidates := make([]string, count)
	for i := range objects {
		rates[i] = rateCalc(actor, objects[i])
	}
	inds := src.Reverse(src.ArgsortNew(rates))[:count]
	for i, ind := range inds {
		candidates[i] = objects[ind].ID
	}
	return candidates, nil
}

func (als *ALS) getActorByID(ctx context.Context, actorID string) (Actor, error) {
	actor, err := als.actorStore.GetActorMFByID(ctx, actorID)
	if errors.Is(err, ErrRecordNotFound) {
		actor = als.newActor(actorID)
		als.actorStore.CreateActorMF(ctx, actor)
	} else if err != nil {
		return Actor{}, werrors.Wrapf(err, "can't get actor by id %s", actorID)
	}
	return actor, nil
}

func (als *ALS) getObjectByID(ctx context.Context, objectID string) (Object, error) {
	object, err := als.objectStore.GetObjectMFByID(ctx, objectID)
	if errors.Is(err, ErrRecordNotFound) {
		object = als.newObject(objectID)
		als.objectStore.CreateObjectMF(ctx, object)
	} else if err != nil {
		return Object{}, werrors.Wrapf(err, "can't get object by id %s", objectID)
	}
	return object, nil
}

func (als *ALS) GetUnitsFromActionPairs(ctx context.Context, actionPairs []recommend.ActionPair) (map[string]Actor, map[string]Object, error) {
	actors := make(map[string]Actor)
	objects := make(map[string]Object)
	for _, action := range actionPairs {
		if _, ok := actors[action.ActorID]; !ok {
			actor, err := als.getActorByID(ctx, action.ActorID)
			actors[action.ActorID] = actor
			if err != nil {
				return nil, nil, err
			}
		}

		if _, ok := objects[action.ObjectID]; !ok {
			object, err := als.getObjectByID(ctx, action.ObjectID)
			objects[action.ObjectID] = object
			if err != nil {
				return nil, nil, err
			}
		}
	}
	return actors, objects, nil
}

func (als *ALS) FitActions(ctx context.Context, actions []recommend.Action) error {
	src.RandomShuffle(&als.rng, actions)
	actionPairs, _ := recommend.SplitActions(actions)
	actors, objects, err := als.GetUnitsFromActionPairs(ctx, actionPairs)
	if err != nil {
		return err
	}

	for _, action := range actions {
		actor := actors[action.ActorID]
		object := objects[action.ObjectID]
		als.gradStep(&actor, &object, action.Rate)
		actors[action.ActorID] = actor
		objects[action.ObjectID] = object
	}

	for _, actor := range actors {
		if err := als.actorStore.UpsertActorMF(ctx, actor); err != nil {
			return werrors.Wrapf(err, "can't upsert actor %s", actor.ID)
		}
	}

	for _, object := range objects {
		if err := als.objectStore.UpsertObjectMF(ctx, object); err != nil {
			return werrors.Wrapf(err, "can't upsert object %s", object.ID)
		}
	}

	return nil
}

func (als *ALS) GetFeatures(ctx context.Context, actionPairs []recommend.ActionPair) ([]recommend.Feature, error) {
	features := make([]recommend.Feature, len(actionPairs))
	for i := range actionPairs {
		actor, err := als.getActorByID(ctx, actionPairs[i].ActorID)
		if err != nil {
			return nil, werrors.Wrap(err, "can't get actor by id")
		}
		object, err := als.getObjectByID(ctx, actionPairs[i].ObjectID)
		if err != nil {
			return nil, werrors.Wrap(err, "can't get object by id")
		}
		features[i].Numeric = []float32{rateCalc(actor, object)}
	}
	return features, nil
}

func (als *ALS) Retrain(ctx context.Context) error {
	actions, err := als.actionStore.ScanAll(ctx)
	if err != nil {
		return err
	}
	err = als.FitActions(ctx, actions)
	if err != nil {
		return werrors.Wrap(err, "can't fit actions")
	}
	return err
}

func (als *ALS) RateCalculate(ctx context.Context, actionPairs []recommend.ActionPair) ([]float32, error) {
	actors, objects, err := als.GetUnitsFromActionPairs(ctx, actionPairs)
	if err != nil {
		return nil, err
	}
	rates := make([]float32, len(actionPairs))
	for i := range actionPairs {
		rates[i] = rateCalc(actors[actionPairs[i].ActorID], objects[actionPairs[i].ObjectID])
	}
	return rates, nil
}
