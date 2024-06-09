package accumulator

import (
	"fmt"
	"math/big"

	"github.com/x-x-x-Ilya/accumulator-worker/pkg/errors"
	syncpkg "github.com/x-x-x-Ilya/accumulator-worker/pkg/sync"
)

type Accumulator struct {
	syncIntVariable syncIntVariable
}

func New() (*Accumulator, error) {
	syncVariable := syncpkg.New()
	if syncVariable == nil {
		return nil, fmt.Errorf("syncIntVariable is nil: %w", errors.ErrInternalServerError)
	}

	return &Accumulator{
		syncIntVariable: syncVariable,
	}, nil
}

func (p *Accumulator) Accumulate(arr []int) {
	var sum int

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	p.syncIntVariable.Add(sum)
}

func (p *Accumulator) Get() big.Int {
	return p.syncIntVariable.Get()
}
