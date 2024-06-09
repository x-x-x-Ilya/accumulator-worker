package worker

import (
	"context"
	"fmt"

	"github.com/x-x-x-Ilya/accumulator-worker/pkg/errors"
)

type Worker struct {
	jobsChan chan func() error
}

func New(ctx context.Context, workersAmount int) (*Worker, error) {
	if workersAmount <= 0 {
		return nil, fmt.Errorf("workersAmount must be positive: %w", errors.ErrBadInput)
	}

	worker := &Worker{
		jobsChan: make(chan func() error, workersAmount),
	}

	for i := 0; i < workersAmount; i++ {
		worker.initWorker(ctx)
	}

	return worker, nil
}

func (w *Worker) initWorker(ctx context.Context) {
	go func() {
		for {
			select {
			case job := <-w.jobsChan:
				if err := job(); err != nil {
					fmt.Println(err)
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (w *Worker) ProcessJob(job func() error) {
	if len(w.jobsChan) == cap(w.jobsChan) {
		fmt.Println("[WARNING] worker input is too large")
	}

	w.jobsChan <- job
}
