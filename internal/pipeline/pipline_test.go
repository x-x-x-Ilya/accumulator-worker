package pipeline_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"golang.org/x/sync/errgroup"

	"github.com/x-x-x-Ilya/accumulator-worker/internal/pipeline"
	accumulatorpkg "github.com/x-x-x-Ilya/accumulator-worker/pkg/accumulator"
	errorspkg "github.com/x-x-x-Ilya/accumulator-worker/pkg/errors"
	processorpkg "github.com/x-x-x-Ilya/accumulator-worker/pkg/processor"
	publisherpkg "github.com/x-x-x-Ilya/accumulator-worker/pkg/publisher"
	workerpkg "github.com/x-x-x-Ilya/accumulator-worker/pkg/worker"
)

//go:generate mockgen -source dependencies.go -destination mock.go -package pipeline
func TestExec(t *testing.T) {
	testCases := []struct {
		name         string
		generatedArr []int
	}{
		{
			name:         "expected with simple array output [0, 24]",
			generatedArr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:         "expected output [0, 24] with reversed generated arr",
			generatedArr: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			name:         "expected output [0, 24] with shuffled arr",
			generatedArr: []int{2, 7, 8, 6, 3, 9, 5, 0, 4, 1},
		},
		{
			name:         "expected output [0, 0] with all zeros arr",
			generatedArr: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:         "expected output [0, -3] with all -1 arr",
			generatedArr: []int{-1, -1, -1, -1, -1, -1, -1 - 1, -1, -1},
		},
	}

	for i := range testCases {
		t.Run(testCases[i].name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())

			ctrl := gomock.NewController(t)

			mockSliceGenerator := pipeline.NewMocksliceGenerator(ctrl)

			publisher := publisherpkg.New()

			accumulator, err := accumulatorpkg.New()
			require.NoError(t, err)

			processor := processorpkg.New()

			worker, err := workerpkg.New(ctx, 10)
			require.NoError(t, err)

			piplineSvc, err := pipeline.New(accumulator, worker, publisher, mockSliceGenerator, processor)
			require.NoError(t, err)

			mockSliceGenerator.EXPECT().MakeRandomInt(10, 10, 1).Return(testCases[i].generatedArr, nil).MaxTimes(5)

			go func() {
				err = piplineSvc.Exec(ctx, 10, 10, 1, time.Second, time.Second-time.Second/10)
			}()

			time.Sleep(time.Second * 2)
			require.NoError(t, err)

			cancel()
		})
	}
}

func TestExecMuLock(t *testing.T) {
	t.Run("mutex error expected", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)

		mockSliceGenerator := pipeline.NewMocksliceGenerator(ctrl)

		publisher := publisherpkg.New()

		accumulator, err := accumulatorpkg.New()
		require.NoError(t, err)

		processor := processorpkg.New()

		worker, err := workerpkg.New(ctx, 10)
		require.NoError(t, err)

		piplineSvc, err := pipeline.New(accumulator, worker, publisher, mockSliceGenerator, processor)
		require.NoError(t, err)

		mockSliceGenerator.EXPECT().MakeRandomInt(10, 10, 1).Return([]int{1, 1, 2, 3, 4, 5, 6, 7, 8, 9}, nil).MaxTimes(5)

		childCtx, cancelChild := context.WithCancel(ctx)
		group, _ := errgroup.WithContext(childCtx)

		var err1 error

		group.Go(
			func() error {
				err1 = piplineSvc.Exec(childCtx, 10, 10, 1, time.Second, time.Second-time.Second/10)

				return err1
			})

		time.Sleep(time.Second)

		var err2 error
		group.Go(
			func() error {
				err2 = piplineSvc.Exec(childCtx, 10, 10, 1, time.Second, time.Second-time.Second/10)

				return err2
			})

		time.Sleep(time.Second)
		fmt.Println("err1 ", err1)
		fmt.Println("err2 ", err2)

		require.ErrorIs(t, err2, errorspkg.ErrInternalServerError)

		cancelChild()
		err = group.Wait()
		require.Error(t, err)
	})
}

func TestExecMuUnlock(t *testing.T) {
	t.Run("mutex error is not expected", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())

		childCtx, cancelChild := context.WithCancel(ctx)
		group, _ := errgroup.WithContext(childCtx)

		ctrl := gomock.NewController(t)

		mockSliceGenerator := pipeline.NewMocksliceGenerator(ctrl)

		publisher := publisherpkg.New()

		accumulator, err := accumulatorpkg.New()
		require.NoError(t, err)

		processor := processorpkg.New()

		worker, err := workerpkg.New(ctx, 10)
		require.NoError(t, err)

		piplineSvc, err := pipeline.New(accumulator, worker, publisher, mockSliceGenerator, processor)
		require.NoError(t, err)

		mockSliceGenerator.EXPECT().MakeRandomInt(10, 10, 1).Return([]int{1, 1, 2, 3, 4, 5, 6, 7, 8, 9}, nil).MaxTimes(5)

		group.Go(
			func() error {
				return piplineSvc.Exec(childCtx, 10, 10, 1, time.Second, time.Second-time.Second/10)
			})

		time.Sleep(time.Second * 1)
		cancelChild()

		time.Sleep(time.Second)

		childCtx2, childCtx2Cancel := context.WithCancel(ctx)
		defer childCtx2Cancel()

		group.Go(
			func() error {
				return piplineSvc.Exec(childCtx2, 10, 10, 1, time.Second, time.Second-time.Second/10)
			})

		time.Sleep(time.Second * 1)
		require.NoError(t, err)

		cancel()

		err = group.Wait()
		require.NoError(t, err)
	})
}
