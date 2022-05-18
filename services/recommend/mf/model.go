package mf

import (
	"errors"

	"github.com/lib/pq"
	"gitlab.com/yars-ai/yars/services/recommend"
)

var ErrRecordNotFound = errors.New("record not found")

type Object struct {
	recommend.Object
	Vector pq.Float32Array `gorm:"type:float[]"`
	Bias   float32
}

type Actor struct {
	recommend.Actor
	Vector pq.Float32Array `gorm:"type:float[]"`
	Bias   float32
}
