package application

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"

	"github.com/x-x-x-Ilya/accumulator-worker/internal/config/env"
	"github.com/x-x-x-Ilya/accumulator-worker/internal/helpers"
	"github.com/x-x-x-Ilya/accumulator-worker/internal/pipeline"
	"github.com/x-x-x-Ilya/accumulator-worker/pkg/accumulator"
	errorspkg "github.com/x-x-x-Ilya/accumulator-worker/pkg/errors"
	"github.com/x-x-x-Ilya/accumulator-worker/pkg/generator"
	"github.com/x-x-x-Ilya/accumulator-worker/pkg/processor"
	"github.com/x-x-x-Ilya/accumulator-worker/pkg/publisher"
	workerpkg "github.com/x-x-x-Ilya/accumulator-worker/pkg/worker"
)

func Start(ctx context.Context, configPath string) error {
	cfg, err := env.NewConfig(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	group, groupCtx := errgroup.WithContext(ctx)

	publisher := publisher.New()

	accumulator, err := accumulator.New()
	if err != nil {
		return fmt.Errorf("failed to create accumulator: %w", err)
	}

	sliceGenerator := generator.NewSliceGenerator()

	processor := processor.New()

	worker, err := workerpkg.New(ctx, cfg.WorkersAmount())
	if err != nil {
		return fmt.Errorf("failed to create worker: %w", err)
	}

	pipelineInstance, err := pipeline.New(accumulator, worker, publisher, sliceGenerator, processor)
	if err != nil {
		return fmt.Errorf("failed to create pipeline: %w", err)
	}

	group.Go(func() error {
		err := pipelineInstance.Exec(groupCtx, cfg.ArrayLength(), cfg.GenerateDelay(), cfg.PrintDelay())
		if err != nil {
			return fmt.Errorf("failed to execute pipeline: %w", err)
		}

		return nil
	})

	group.Go(func() error {
		err := helpers.ListenSystemSignals(groupCtx)
		if err != nil {
			return fmt.Errorf("system signal received: %w", err)
		}

		return nil
	})

	err = group.Wait()
	if err != nil {
		if errors.Is(err, errorspkg.ErrApplicationShutdown) {
			fmt.Println("app.Start() graceful shutdown succeeded")
			return nil
		}

		return fmt.Errorf("graceful shutdown failed: %w", err)
	}

	return nil
}
