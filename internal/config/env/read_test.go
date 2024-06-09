package env

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMustPositiveDuration(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		isError  bool
		expected time.Duration
	}{
		{
			name:     "simple case",
			input:    "1s",
			expected: time.Second,
		},
		{
			name:    "input is negative",
			input:   "-1s",
			isError: true,
		},
		{
			name:    "input is invalid (string)",
			input:   "sdfoijosjdfo",
			isError: true,
		},
		{
			name:    "input is invalid (spaces)",
			input:   "   3s  ",
			isError: true,
		},
		{
			name:    "input is invalid (empty string)",
			input:   "",
			isError: true,
		},
		{
			name:    "input is invalid (int)",
			input:   "1",
			isError: true,
		},
	}

	for i := range testCases {
		t.Run(testCases[i].name, func(t *testing.T) {
			err := os.Setenv("test_key", testCases[i].input)
			require.NoError(t, err)

			actual, err := mustPositiveDuration("test_key")
			if testCases[i].isError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				require.Equal(t, testCases[i].expected, actual)
			}
		})
	}
}

func TestMustPositiveDurationWithoutEnv(t *testing.T) {
	t.Run("no env", func(t *testing.T) {
		os.Clearenv()

		_, err := mustPositiveDuration("test_key")
		require.Error(t, err)
	})
}

func TestMustPositiveInt(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		isError  bool
		expected int
	}{
		{
			name:     "simple case",
			input:    "1",
			expected: 1,
		},
		{
			name:    "input is negative",
			input:   "-1",
			isError: true,
		},
		{
			name:    "input is invalid (string)",
			input:   "sdfoijosjdfo",
			isError: true,
		},
		{
			name:    "input is invalid (spaces)",
			input:   "   3  ",
			isError: true,
		},
		{
			name:    "input is invalid (empty string)",
			input:   "",
			isError: true,
		},
		{
			name:    "input is invalid (duration)",
			input:   "1s",
			isError: true,
		},
	}

	for i := range testCases {
		t.Run(testCases[i].name, func(t *testing.T) {
			err := os.Setenv("test_key", testCases[i].input)
			require.NoError(t, err)

			actual, err := mustPositiveInt("test_key")
			if testCases[i].isError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				require.Equal(t, testCases[i].expected, actual)
			}
		})
	}
}

func TestMustPositiveIntWithoutEnv(t *testing.T) {
	t.Run("no env", func(t *testing.T) {
		os.Clearenv()

		_, err := mustPositiveInt("test_key")
		require.Error(t, err)
	})
}

func TestShouldPositiveInt(t *testing.T) {
	testCases := []struct {
		name       string
		input      string
		defaultVal int
		isError    bool
		expected   int
	}{
		{
			name:     "input is positive (default is zero)",
			input:    "1",
			expected: 1,
		},
		{
			name:       "input is negative (default is zero)",
			input:      "-1",
			defaultVal: 0,
			isError:    true,
		},
		{
			name:       "input is invalid (default is negative)",
			input:      "sdfoijosjdfo",
			defaultVal: -5,
			isError:    true,
		},
		{
			name:       "input is invalid (default is positive)",
			input:      "   3  ",
			defaultVal: 5,
			isError:    true,
		},
		{
			name:       "input is empty (default is positive)",
			input:      "",
			defaultVal: 5,
			isError:    true,
		},
	}

	for i := range testCases {
		t.Run(testCases[i].name, func(t *testing.T) {
			err := os.Setenv("test_key", testCases[i].input)
			require.NoError(t, err)

			actual, err := shouldPositiveInt("test_key", testCases[i].defaultVal)
			if testCases[i].isError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				require.Equal(t, testCases[i].expected, actual)
			}
		})
	}
}

func TestShouldPositiveIntWithDefault(t *testing.T) {
	testCases := []struct {
		name       string
		defaultVal int
		isError    bool
		expected   int
	}{
		{
			name:       "no input (default is positive)",
			defaultVal: 5,
			expected:   5,
		},
		{
			name:       "no input (default is zero)",
			defaultVal: 0,
			isError:    true,
		},
		{
			name:       "no input (default is negative)",
			defaultVal: -1,
			isError:    true,
		},
	}

	for i := range testCases {
		t.Run(testCases[i].name, func(t *testing.T) {
			os.Clearenv()
			actual, err := shouldPositiveInt("test_key", testCases[i].defaultVal)
			if testCases[i].isError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				require.Equal(t, testCases[i].expected, actual)
			}
		})
	}
}
