package pipeline

import (
	"math/big"
)

type accumulator interface {
	Accumulate(input []int)
	Get() big.Int
}

type publisher interface {
	Publish(data any)
}

type worker interface {
	ProcessJob(job func() error)
}

type processor interface {
	ThreeMaxElements(input []int) ([]int, error)
}

type sliceGenerator interface {
	MakeRandomInt(length int) ([]int, error)
}
