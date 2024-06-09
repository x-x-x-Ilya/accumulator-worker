package pipeline

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/x-x-x-Ilya/accumulator-worker/internal/helpers"
	errorspkg "github.com/x-x-x-Ilya/accumulator-worker/pkg/errors"
)

type Pipeline struct {
	accumulator    accumulator
	worker         worker
	publisher      publisher
	sliceGenerator sliceGenerator
	processor      processor

	mu *sync.Mutex
}

func New(
	accumulator accumulator,
	worker worker,
	publisher publisher,
	sliceGenerator sliceGenerator,
	processor processor,
) (*Pipeline, error) {
	if accumulator == nil {
		return nil, fmt.Errorf("accumulator is nil: %w", errorspkg.ErrBadInput)
	}

	if worker == nil {
		return nil, fmt.Errorf("worker is nil: %w", errorspkg.ErrBadInput)
	}

	if publisher == nil {
		return nil, fmt.Errorf("publisher is nil: %w", errorspkg.ErrBadInput)
	}

	if sliceGenerator == nil {
		return nil, fmt.Errorf("sliceGenerator is nil: %w", errorspkg.ErrBadInput)
	}

	if processor == nil {
		return nil, fmt.Errorf("processor is nil: %w", errorspkg.ErrBadInput)
	}

	return &Pipeline{
		accumulator:    accumulator,
		publisher:      publisher,
		worker:         worker,
		sliceGenerator: sliceGenerator,
		processor:      processor,
		mu:             &sync.Mutex{},
	}, nil
}

func (p *Pipeline) validateExecInput(arrayLength int, generateDelay, printDelay time.Duration) error {
	if arrayLength < 3 {
		return fmt.Errorf("arrayLength must be 3 or greater: %w", errorspkg.ErrBadInput)
	}

	if generateDelay <= 0 {
		return fmt.Errorf("generateDelay must be positive: %w", errorspkg.ErrBadInput)
	}

	if printDelay <= 0 {
		return fmt.Errorf("printDelay must be positive: %w", errorspkg.ErrBadInput)
	}

	return nil
}

func (p *Pipeline) Exec(ctx context.Context, arrayLength int, generateDelay, printDelay time.Duration) error {
	if err := p.validateExecInput(arrayLength, generateDelay, printDelay); err != nil {
		return fmt.Errorf("validateExecInput failed: %w", err)
	}

	if !p.mu.TryLock() {
		return fmt.Errorf("pipline is already running: %w", errorspkg.ErrInternalServerError)
	}
	defer func() {
		p.mu.Unlock()
	}()

	group, ctx := errgroup.WithContext(ctx)

	var (
		randomArrChan        = make(chan []int, 1)
		threeMaxElementsChan = make(chan []int, 3)
	)

	group.Go(
		func() error {
			return p.generatorWorker(ctx, generateDelay, arrayLength, randomArrChan)
		},
	)

	group.Go(
		func() error {
			return p.processorWorker(ctx, randomArrChan, threeMaxElementsChan)
		},
	)

	group.Go(
		func() error {
			return p.accumulatorWorker(ctx, threeMaxElementsChan)
		},
	)

	group.Go(
		func() error {
			return p.publisherWorker(ctx, printDelay)
		},
	)

	fmt.Println("Exec() started successfully")

	err := group.Wait()
	if err != nil {
		if errors.Is(err, errorspkg.ErrApplicationShutdown) {
			fmt.Println("Exec() graceful shutdown succeeded")
			return nil
		}

		return fmt.Errorf("exec() graceful shutdown failed: %w", err)
	}

	return nil
}

func (p *Pipeline) generatorWorker(ctx context.Context, generateDelay time.Duration, arrayLength int, outputChan chan []int) error {
	generateTicker := time.NewTicker(generateDelay)

	for {
		select {
		case <-generateTicker.C:
			randomArr, err := p.sliceGenerator.MakeRandomInt(arrayLength)
			if err != nil {
				generateTicker.Stop()
				return fmt.Errorf("generate random slice generator failed: %w", err)
			}

			if err := helpers.WriteToChan(ctx, randomArr, outputChan); err != nil {
				generateTicker.Stop()
				return fmt.Errorf("error pushing random array: %w", err)
			}

		case <-ctx.Done():
			generateTicker.Stop()
			return fmt.Errorf("generate job finished ctx.Done: %w", errorspkg.ErrApplicationShutdown)
		}
	}
}

func (p *Pipeline) processorWorker(ctx context.Context, inputArrChan, outputArrChan chan []int) error {
	for {
		randomArr, err := helpers.ReadFromChan(ctx, inputArrChan)
		if err != nil {
			return fmt.Errorf("error reading: %w", err)
		}

		go p.worker.ProcessJob(
			func() error {
				arr, err := p.processor.ThreeMaxElements(randomArr)
				if err != nil {
					return fmt.Errorf("three max elements calculator error: %w", err)
				}

				err = helpers.WriteToChan(ctx, arr, outputArrChan)
				if err != nil {
					return fmt.Errorf("sending max elements to chan error: %w", err)
				}

				return nil
			})
	}
}

func (p *Pipeline) accumulatorWorker(ctx context.Context, inputChan chan []int) error {
	for {
		threeMaxElements, err := helpers.ReadFromChan(ctx, inputChan)
		if err != nil {
			return fmt.Errorf("error reading max elements from channel: %w", err)
		}

		p.accumulator.Accumulate(threeMaxElements)
	}
}

func (p *Pipeline) publisherWorker(ctx context.Context, printDelay time.Duration) error {
	printTicker := time.NewTicker(printDelay)

	for {
		select {
		case <-printTicker.C:
			currentAccumulator := p.accumulator.Get()

			p.publisher.Publish(currentAccumulator.String())
		case <-ctx.Done():
			printTicker.Stop()
			return fmt.Errorf("generate job finished ctx.Done: %w", errorspkg.ErrApplicationShutdown)
		}
	}
}
