package sync

import (
	"math/big"
	"sync"
)

type BigSync struct {
	*sync.RWMutex
	value big.Int
}

func New() *BigSync {
	mu := &sync.RWMutex{}

	return &BigSync{
		RWMutex: mu,
	}
}

func (g *BigSync) Add(newValue int) {
	g.Lock()
	defer g.Unlock()

	g.value.Add(&g.value, big.NewInt(int64(newValue)))
}

func (g *BigSync) Get() big.Int {
	g.RLock()
	defer g.RUnlock()

	return g.value
}
