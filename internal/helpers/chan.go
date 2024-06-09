package helpers

import (
	"context"

	"github.com/x-x-x-Ilya/accumulator-worker/pkg/errors"
)

func WriteToChan[T any](ctx context.Context, val []T, ch chan []T) error {
	select {
	case ch <- val:
		return nil
	case <-ctx.Done():
		return errors.ErrApplicationShutdown
	}
}

func ReadFromChan[T any](ctx context.Context, ch chan []T) ([]T, error) {
	select {
	case val := <-ch:
		return val, nil
	case <-ctx.Done():
		return nil, errors.ErrApplicationShutdown
	}
}
