package src

import (
	"math"
	"sort"
)

func Dot(a, b []float32) float32 {
	if len(a) != len(b) {
		panic("floats: slice lengths do not match")
	}
	var ret float32 = 0
	for i := range a {
		ret += a[i] * b[i]
	}
	return ret
}

func MulVV(a, b []float32) []float32 {
	if len(a) != len(b) {
		panic("floats: slice lengths do not match")
	}
	ret := make([]float32, len(a))
	for i := range a {
		ret[i] = a[i] * b[i]
	}
	return ret
}

func MulVS(a []float32, b float32) []float32 {
	ret := make([]float32, len(a))
	for i := range a {
		ret[i] = a[i] * b
	}
	return ret
}

func AddVV(a, b []float32) []float32 {
	if len(a) != len(b) {
		panic("floats: slice lengths do not match")
	}
	ret := make([]float32, len(a))
	for i := range a {
		ret[i] = a[i] + b[i]
	}
	return ret
}

func SubVV(a, b []float32) []float32 {
	if len(a) != len(b) {
		panic("floats: slice lengths do not match")
	}
	ret := make([]float32, len(a))
	for i := range a {
		ret[i] = a[i] - b[i]
	}
	return ret
}

func Norm(a []float32) []float32 {
	var sum float32 = 0
	for _, val := range a {
		sum += val * val
	}
	sum = 1 / float32(math.Sqrt(float64(sum)))
	ret := make([]float32, len(a))
	for i := range a {
		ret[i] = a[i] * sum
	}
	return ret
}

func NormIf(a []float32) []float32 {
	var sum float32 = 0
	for _, val := range a {
		sum += val * val
	}
	sum = 1 / float32(math.Sqrt(float64(sum)))
	if sum > 1 {
		sum = 1
	}
	ret := make([]float32, len(a))
	for i := range a {
		ret[i] = a[i] * sum
	}
	return ret
}

type argsort struct {
	s    []float32
	inds []int
}

func (a argsort) Len() int {
	return len(a.s)
}

func (a argsort) Less(i, j int) bool {
	return a.s[a.inds[i]] < a.s[a.inds[j]]
}

func (a argsort) Swap(i, j int) {
	a.inds[i], a.inds[j] = a.inds[j], a.inds[i]
}

func ArgsortNew(src []float32) []int {
	inds := make([]int, len(src))
	for i := range src {
		inds[i] = i
	}
	Argsort(src, inds)
	return inds
}

func Argsort(src []float32, inds []int) {
	if len(src) != len(inds) {
		panic("floats: length of inds does not match length of slice")
	}
	a := argsort{s: src, inds: inds}
	sort.Sort(a)
}

func Reverse[E any](src []E) []E {
	ret := make([]E, len(src))
	copy(ret, src)
	for i := 0; 2*i < len(src); i++ {
		j := len(src) - i - 1
		ret[j], ret[i] = ret[i], ret[j]
	}
	return ret
}

func MinInt32(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
