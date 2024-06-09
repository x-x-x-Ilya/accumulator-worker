package env

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewConfig(t *testing.T) {
	expected := Config{
		printer: printer{
			printDelay: time.Second,
		},
		processor: processor{
			workersAmount: 3,
		},
		arrayGenerator: arrayGenerator{
			generatorDelay: time.Millisecond * 2,
			arrayLength:    10,
		},
	}

	cfg, err := NewConfig("env-template")
	require.NoError(t, err)

	require.Equal(t, expected, *cfg)

	require.Equal(t, expected.arrayLength, cfg.ArrayLength())
	require.Equal(t, expected.printDelay, cfg.PrintDelay())
	require.Equal(t, expected.generatorDelay, cfg.GenerateDelay())
	require.Equal(t, expected.workersAmount, cfg.WorkersAmount())
}
