package application_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/x-x-x-Ilya/accumulator-worker/internal/application"
)

func TestStart(t *testing.T) {
	t.Run("start test", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		childCtx, childCancel := context.WithCancel(ctx)
		group, groupCtx := errgroup.WithContext(childCtx)

		var err error
		group.Go(
			func() error {
				err = application.Start(groupCtx, "../../.env")
				if err != nil {
					return fmt.Errorf("start error: %w", err)
				}

				return nil
			})

		time.Sleep(time.Second * 1)
		childCancel()

		groupErr := group.Wait()
		require.NoError(t, groupErr)
		require.NoError(t, err)
	})
}
