package processor

import (
	"fmt"
	"math"

	"github.com/x-x-x-Ilya/accumulator-worker/pkg/errors"
)

type Processor struct{}

func New() *Processor {
	return &Processor{}
}

func (p *Processor) ThreeMaxElements(input []int) ([]int, error) {
	if len(input) < 3 {
		return nil, fmt.Errorf("input should be at least 3 elements: %w", errors.ErrBadInput)
	}

	arr := make([]int, 3)

	for i := 0; i < len(arr); i++ {
		arr[i] = math.MinInt
	}

	for i := 0; i < len(input); i++ {
		for j := range arr {
			if arr[j] < input[i] {
				p.update(arr, input[i], j)
				break
			}
		}
	}

	return arr, nil
}

func (p *Processor) update(maxElements []int, newValue int, index int) {
	for i := index; i < len(maxElements); i++ {
		newValue, maxElements[i] = maxElements[i], newValue
	}
}
