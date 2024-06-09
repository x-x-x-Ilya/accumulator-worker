package accumulator

import (
	"math/big"
)

type syncIntVariable interface {
	Get() big.Int
	Add(newValue int)
}
