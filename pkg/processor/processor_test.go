package processor_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/x-x-x-Ilya/accumulator-worker/pkg/processor"
)

func TestThreeMax(t *testing.T) {
	processorSvc := processor.New()

	testCases := []struct {
		name     string
		input    []int
		isError  bool
		expected []int
	}{
		{
			name:     "simple case",
			input:    []int{7, 2, 5, 3, 1, 2, 5, 4},
			expected: []int{7, 5, 5},
		},
		{
			name:     "simple case 2",
			input:    []int{7, 2, 5, 3, 11, 2, 5, 4},
			expected: []int{11, 7, 5},
		},
		{
			name:    "0",
			input:   []int{0},
			isError: true,
		},
		{
			name:     "0 1 0",
			input:    []int{0, 1, 0},
			expected: []int{1, 0, 0},
		},
		{
			name:    "empty array",
			input:   []int{},
			isError: true,
		},
		{
			name:     "positive and negative array",
			input:    []int{-1, 0, -8, 9},
			expected: []int{9, 0, -1},
		},
		{
			name:     "negative array",
			input:    []int{-1, -8, -6},
			expected: []int{-1, -6, -8},
		},
	}

	for i := range testCases {
		t.Run(testCases[i].name, func(t *testing.T) {
			actual, err := processorSvc.ThreeMaxElements(testCases[i].input)

			if testCases[i].isError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				require.Equal(t, testCases[i].expected, actual)
			}
		})
	}
}
