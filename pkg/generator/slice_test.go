package generator_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/x-x-x-Ilya/accumulator-worker/pkg/generator"
)

func TestGenerate(t *testing.T) {
	sliceGenerator := generator.NewSliceGenerator()

	testCases := []struct {
		name            string
		input           int
		isErrorExpected bool
	}{
		{
			name:  "simple case",
			input: 5,
		},
		{
			name:  "0",
			input: 0,
		},
		{
			name:            "-1",
			input:           -1,
			isErrorExpected: true,
		},
	}

	for i := range testCases {
		t.Run(testCases[i].name, func(t *testing.T) {
			generatedArr, err := sliceGenerator.MakeRandomInt(testCases[i].input, 100, 1)

			if testCases[i].isErrorExpected {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, len(generatedArr), testCases[i].input)
			}
		})
	}
}
