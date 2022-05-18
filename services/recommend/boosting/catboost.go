package boosting

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"

	werrors "github.com/pkg/errors"
	"gitlab.com/yars-ai/yars/catboost"
	"gitlab.com/yars-ai/yars/services/recommend"
	"gitlab.com/yars-ai/yars/src"
)

func NewCatboost(
	name string,
	workDir string,
	actionStore recommend.SquashActionStore,
	featureGenerator recommend.FeatureGenerator,
	candidateGenerator recommend.CanditateGenerator,
) *Catboost {
	var model *catboost.Model = nil
	if _, err := os.Stat(path.Join(workDir, "train/model.bin")); !errors.Is(err, os.ErrNotExist) {
		model, _ = catboost.LoadFullModelFromFile(path.Join(workDir, "train/model.bin"))
	}
	return &Catboost{
		name:               name,
		workDir:            workDir,
		actionStore:        actionStore,
		featureGenerator:   featureGenerator,
		candidateGenerator: candidateGenerator,
		model:              model,
	}
}

type Catboost struct {
	name               string
	workDir            string
	actionStore        recommend.SquashActionStore
	featureGenerator   recommend.FeatureGenerator
	candidateGenerator recommend.CanditateGenerator
	model              *catboost.Model
}

func (cb *Catboost) Retrain(ctx context.Context) error {
	actions, err := cb.actionStore.ScanAll(ctx)
	if err != nil {
		return werrors.Wrap(err, "can't scan all actions")
	}
	actionPairs, predictiveLabels := recommend.SplitActions(actions)
	features, err := cb.featureGenerator.GetFeatures(ctx, actionPairs)
	if err != nil {
		return err
	}

	if err := recommend.SaveFeatureToFile(features, predictiveLabels, cb.workDir); err != nil {
		return err
	}

	params := []string{
		"fit",
		fmt.Sprintf("--learn-set=%s", path.Join(cb.workDir, "train.csv")),
		fmt.Sprintf("--test-set=%s", path.Join(cb.workDir, "test.csv")),
		fmt.Sprintf("--cd=%s", path.Join(cb.workDir, "info.cd")),
		"--delimiter=,",
		"--loss-function=RMSE",
		"--model-format=CatboostBinary",
		fmt.Sprintf("--train-dir=%s", path.Join(cb.workDir, "train")),
	}

	cmd := exec.Command("./catboost-linux", params...)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err = cmd.Run()
	if err != nil {
		return werrors.Wrapf(err, "err %s", errb.String())
	}
	fmt.Printf("out:%s", outb.String())

	cb.model, err = catboost.LoadFullModelFromFile(path.Join(cb.workDir, "train/model.bin"))
	if err != nil {
		return werrors.Wrap(err, "can't load catboost model from file")
	}

	return nil
}

func (cb *Catboost) RateCalculate(ctx context.Context, actionPairs []recommend.ActionPair) ([]float32, error) {
	features, err := cb.featureGenerator.GetFeatures(ctx, actionPairs)
	if err != nil {
		return nil, werrors.Wrap(err, "can't get features for catboost")
	}
	numFeatures := make([][]float32, len(features))
	catFeatures := make([][]string, len(features))
	for i := range features {
		numFeatures[i] = features[i].Numeric
		catFeatures[i] = features[i].Categorial
	}
	res64, err := cb.model.CalcModelPrediction(numFeatures, len(numFeatures[0]), catFeatures, len(catFeatures[0]))
	if err != nil {
		return nil, werrors.Wrap(err, "can't calculate model prediction")
	}
	res32 := make([]float32, len(res64))
	for i := range res64 {
		res32[i] = float32(res64[i])
	}
	return res32, nil
}

func (cb *Catboost) GetCanditates(ctx context.Context, actorID string, count int) ([]string, error) {
	candidates, err := cb.candidateGenerator.GetCanditates(ctx, actorID, 5*count)
	if err != nil {
		return nil, err
	}
	actionPairs := recommend.CandidatesToActionPairs(actorID, candidates)
	rates, err := cb.RateCalculate(ctx, actionPairs)
	if err != nil {
		return nil, err
	}
	resIndex := src.Reverse(src.ArgsortNew(rates))[:count]
	answer := make([]string, count)
	for i := range resIndex {
		answer[i] = actionPairs[resIndex[i]].ObjectID
	}
	return answer, nil
}
