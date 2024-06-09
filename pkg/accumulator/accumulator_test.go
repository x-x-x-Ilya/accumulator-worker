package accumulator_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/x-x-x-Ilya/accumulator-worker/pkg/accumulator"
)

func TestAccumulate(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected *big.Int
	}{
		{
			name:     "simple case",
			input:    []int{7, 2, 5, 3, 1, 2, 5, 4},
			expected: big.NewInt(int64(29)),
		},
		{
			name:     "0",
			input:    []int{0},
			expected: big.NewInt(int64(0)),
		},
		{
			name:     "0 1 0",
			input:    []int{0, 1, 0},
			expected: big.NewInt(int64(1)),
		},
		{
			name:     "empty array",
			input:    []int{},
			expected: big.NewInt(int64(0)),
		},
		{
			name:     "positive and negative array",
			input:    []int{-1, 0, -8, 9},
			expected: big.NewInt(int64(0)),
		},
		{
			name:     "negative array",
			input:    []int{-1, -8, -6},
			expected: big.NewInt(int64(-15)),
		},
	}

	for i := range testCases {
		t.Run(testCases[i].name, func(t *testing.T) {
			accumulatorSvc, err := accumulator.New()
			require.NoError(t, err)

			accumulatorSvc.Accumulate(testCases[i].input)

			require.Equal(t, *testCases[i].expected, accumulatorSvc.Get())
		})
	}
}

func TestAccumulates(t *testing.T) {
	testCases := []struct {
		name     string
		input    [][]int
		expected *big.Int
	}{
		{
			name: "simple case",
			input: [][]int{
				{7, 2, 5, 3, 1, 2, 5, 4},
				{-1, -8, -6},
			},
			expected: big.NewInt(int64(14)),
		},
	}

	for i := range testCases {
		t.Run(testCases[i].name, func(t *testing.T) {
			accumulatorSvc, err := accumulator.New()
			require.NoError(t, err)

			for _, input := range testCases[i].input {
				accumulatorSvc.Accumulate(input)
			}

			require.Equal(t, *testCases[i].expected, accumulatorSvc.Get())
		})
	}
}
