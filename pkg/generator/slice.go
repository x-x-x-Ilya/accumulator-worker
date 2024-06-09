package generator

import (
	"math/rand"
	"time"

	"github.com/x-x-x-Ilya/accumulator-worker/pkg/errors"
)

type SliceGenerator struct{}

func NewSliceGenerator() *SliceGenerator {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	return &SliceGenerator{}
}

func (g *SliceGenerator) MakeRandomInt(length, maxValue, minValue int) ([]int, error) {
	if length < 0 {
		return nil, errors.ErrBadInput
	}

	randomInt := make([]int, length)
	for i := 0; i < length; i++ {
		randomInt[i] = rand.Intn(maxValue-minValue+1) + minValue
	}

	return randomInt, nil
}
