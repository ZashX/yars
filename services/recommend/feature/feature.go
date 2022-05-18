package feature

import (
	"context"

	werrors "github.com/pkg/errors"
	"gitlab.com/yars-ai/yars/services/recommend"
)

type UnitFeatureGenerator struct {
	actorUnit  recommend.UnitStoreReader
	objectUnit recommend.UnitStoreReader
}

func NewUnitFeatureGenerator(
	actorUnit recommend.UnitStoreReader,
	objectUnit recommend.UnitStoreReader,
) *UnitFeatureGenerator {
	return &UnitFeatureGenerator{
		actorUnit:  actorUnit,
		objectUnit: objectUnit,
	}
}

func (ufg *UnitFeatureGenerator) GetFeatures(ctx context.Context, actionPairs []recommend.ActionPair) ([]recommend.Feature, error) {
	actorFeatures := make([]recommend.Feature, len(actionPairs))
	objectFeatures := make([]recommend.Feature, len(actionPairs))
	for i := range actionPairs {
		actor, err := ufg.actorUnit.GetUnitById(ctx, actionPairs[i].ActorID)
		if err != nil {
			return nil, werrors.Wrap(err, "can't get actor unit by id")
		}
		actorFeatures[i] = recommend.ExtractFeatures(actor)

		object, err := ufg.objectUnit.GetUnitById(ctx, actionPairs[i].ObjectID)
		if err != nil {
			return nil, werrors.Wrap(err, "can't get object unit by id")
		}
		objectFeatures[i] = recommend.ExtractFeatures(object)
	}
	return recommend.UnionSliceFeatures([][]recommend.Feature{actorFeatures, objectFeatures})
}

type MergedFeatureGenerator struct {
	featureGenerators []recommend.FeatureGenerator
}

func NewMergedFeatureGenerator(featureGenerators []recommend.FeatureGenerator) *MergedFeatureGenerator {
	return &MergedFeatureGenerator{
		featureGenerators: featureGenerators,
	}
}

func (mfg *MergedFeatureGenerator) GetFeatures(ctx context.Context, actionPairs []recommend.ActionPair) ([]recommend.Feature, error) {
	f := make([][]recommend.Feature, len(mfg.featureGenerators))
	for i := range mfg.featureGenerators {
		cur, err := mfg.featureGenerators[i].GetFeatures(ctx, actionPairs)
		if err != nil {
			return nil, err
		}
		f[i] = cur
	}
	return recommend.UnionSliceFeatures(f)
}
