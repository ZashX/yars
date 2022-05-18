package recommend

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"math"
	"os"
	"path"
	"reflect"

	"github.com/lib/pq"
	werrors "github.com/pkg/errors"
)

var ErrFeaturesDifferentLengths = errors.New("features of different lengths")

type Feature struct {
	Numeric    []float32
	Categorial []string
}

type ActionPair struct {
	ActorID  string
	ObjectID string
}

type PredictiveLabel struct {
	Value float32
}

func SplitActions(actions []Action) ([]ActionPair, []PredictiveLabel) {
	actionPairs := make([]ActionPair, len(actions))
	predictiveLabels := make([]PredictiveLabel, len(actions))
	for i := range actions {
		actionPairs[i].ActorID = actions[i].ActorID
		actionPairs[i].ObjectID = actions[i].ObjectID
		predictiveLabels[i].Value = actions[i].Rate
	}
	return actionPairs, predictiveLabels
}

func ExtractFeatures(t interface{}) Feature {
	val := reflect.ValueOf(t).Elem()
	feature := Feature{}
	for i := 0; i < val.NumField(); i++ {
		fieldInterface := val.Field(i).Interface()
		switch field := fieldInterface.(type) {
		case int32:
			feature.Numeric = append(feature.Numeric, float32(field))
		case float32:
			feature.Numeric = append(feature.Numeric, field)
		case pq.Float32Array:
			feature.Numeric = append(feature.Numeric, field...)
		case string:
			feature.Categorial = append(feature.Categorial, field)
		}
	}
	return feature
}

func UnionSliceFeatures(sliceFeatures [][]Feature) ([]Feature, error) {
	n := len(sliceFeatures[0])
	for i := range sliceFeatures {
		if n != len(sliceFeatures[i]) {
			return nil, ErrFeaturesDifferentLengths
		}
	}

	resFeatures := make([]Feature, n)
	for i := range sliceFeatures {
		for j := range sliceFeatures[i] {
			resFeatures[j].Numeric = append(resFeatures[j].Numeric, sliceFeatures[i][j].Numeric...)
			resFeatures[j].Categorial = append(resFeatures[j].Categorial, sliceFeatures[i][j].Categorial...)
		}
	}

	return resFeatures, nil
}

func RMSE(ctx context.Context, rc RateCalculator, actions []Action) (float32, error) {
	actionPairs, _ := SplitActions(actions)
	rate, err := rc.RateCalculate(ctx, actionPairs)
	if err != nil {
		return 0, err
	}
	rmse := float32(0)
	for i := range rate {
		diff := rate[i] - actions[i].Rate
		rmse += diff * diff
	}
	rmse /= float32(len(actions))
	return float32(math.Sqrt(float64(rmse))), nil
}

// func NDCG(ctx context.Context, rc RateCalculator, actions []Action) (float32, error) {
// 	actionPairs, rate := SplitActions(actions)
// 	predict, err := rc.RateCalculate(ctx, actionPairs)
// 	if err != nil {
// 		return 0, err
// 	}

// 	mRate := make(map[string][]float32)
// 	mPredict := make(map[string][]float32)
// 	for i := range actionPairs {
// 		mRate[actionPairs[i].ActorID] = append(mRate[actionPairs[i].ActorID], rate[i].Value)
// 		mPredict[actionPairs[i].ObjectID] = append(mPredict[actionPairs[i].ObjectID], predict[i])
// 	}

// 	for actorID := range mRate {
// 		ratePermutation := src.Reverse(src.ArgsortNew(mRate[actorID]))
// 		predictPermutation := src.Reverse(src.ArgsortNew(mPredict[actorID]))

// 	}

// 	return float32(math.Sqrt(float64(rmse))), nil
// }

func SaveFeatureToFile(features []Feature, predictiveLabels []PredictiveLabel, workDir string) error {
	csvFileTrain, err := os.Create(path.Join(workDir, "train.csv"))
	if err != nil {
		return werrors.Wrap(err, "can't create csv file")
	}
	defer csvFileTrain.Close()
	csvWriterTrain := csv.NewWriter(csvFileTrain)
	defer csvWriterTrain.Flush()

	csvFileTest, err := os.Create(path.Join(workDir, "test.csv"))
	if err != nil {
		return werrors.Wrap(err, "can't create csv file")
	}
	defer csvFileTest.Close()
	csvWriterTest := csv.NewWriter(csvFileTest)
	defer csvWriterTest.Flush()

	for i := range features {
		row := []string{}
		for j := range features[i].Numeric {
			row = append(row, fmt.Sprintf("%f", features[i].Numeric[j]))
		}
		row = append(row, features[i].Categorial...)
		row = append(row, fmt.Sprintf("%f", predictiveLabels[i].Value))
		if i%10 < 2 {
			csvWriterTest.Write(row)
		} else {
			csvWriterTrain.Write(row)
		}
	}
	cdFile, err := os.Create(path.Join(workDir, "info.cd"))
	if err != nil {
		return werrors.Wrap(err, "can't create cd file")
	}
	defer cdFile.Close()

	sizeNumeric, sizeFeature := len(features[0].Numeric), len(features[0].Categorial)
	for i := 0; i < sizeNumeric+sizeFeature; i++ {
		if i < sizeNumeric {
			cdFile.WriteString(fmt.Sprintf("%d\tNum\n", i))
		} else {
			cdFile.WriteString(fmt.Sprintf("%d\tCateg\n", i))
		}
	}
	cdFile.WriteString(fmt.Sprintf("%d\tLabel\n", sizeNumeric+sizeFeature))
	return nil
}

func CandidatesToActionPairs(actorID string, objectIDs []string) []ActionPair {
	actionPairs := make([]ActionPair, len(objectIDs))
	for i := range objectIDs {
		actionPairs[i].ActorID = actorID
		actionPairs[i].ObjectID = objectIDs[i]
	}
	return actionPairs
}
