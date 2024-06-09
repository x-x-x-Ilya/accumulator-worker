package helpers

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/x-x-x-Ilya/accumulator-worker/pkg/errors"
)

func ListenSystemSignals(ctx context.Context) error {
	sigs := []os.Signal{syscall.SIGINT, syscall.SIGTERM}

	ch := make(chan os.Signal, len(sigs))
	signal.Notify(ch, sigs...)
	for {
		select {
		case sig := <-ch:
			fmt.Printf("\r\nterminated with sig %q \r\n", sig)
			return errors.ErrApplicationShutdown
		case <-ctx.Done():
			return errors.ErrApplicationShutdown
		}
	}
}
