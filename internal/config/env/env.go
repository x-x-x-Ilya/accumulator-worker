package env

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"

	"github.com/x-x-x-Ilya/accumulator-worker/pkg/errors"
)

func readArrayGenerator() (arrayGenerator, error) {
	generatorDelay, err := mustPositiveDuration("ARRAY_GENERATOR_DELAY")
	if err != nil {
		return arrayGenerator{}, fmt.Errorf("ARRAY_GENERATOR_DELAY error: %w", err)
	}

	arrayLength, err := shouldPositiveInt("ARRAY_GENERATOR_LENGTH", 10)
	if err != nil {
		return arrayGenerator{}, fmt.Errorf("ARRAY_GENERATOR_LENGTH error: %w", err)
	}

	maxElemValue, err := shouldPositiveInt("ARRAY_GENERATOR_MAX_ELEM_VALUE", 10)
	if err != nil {
		return arrayGenerator{}, fmt.Errorf("ARRAY_GENERATOR_MAX_ELEM_VALUE error: %w", err)
	}

	minElemValue, err := shouldPositiveInt("ARRAY_GENERATOR_MIN_ELEM_VALUE", 1)
	if err != nil {
		return arrayGenerator{}, fmt.Errorf("ARRAY_GENERATOR_MIN_ELEM_VALUE error: %w", err)
	}

	return arrayGenerator{
		generatorDelay: generatorDelay,
		arrayLength:    arrayLength,
		maxElemValue:   maxElemValue,
		minElemValue:   minElemValue,
	}, nil
}

func readPrinter() (printer, error) {
	printDelay, err := mustPositiveDuration("PRINT_DELAY")
	if err != nil {
		return printer{}, fmt.Errorf("PRINT_DELAY error: %w", err)
	}

	return printer{
		printDelay: printDelay,
	}, nil
}

func readProcessor() (processor, error) {
	workersAmount, err := mustPositiveInt("WORKERS_AMOUNT")
	if err != nil {
		return processor{}, fmt.Errorf("WORKERS_AMOUNT error: %w", err)
	}

	return processor{
		workersAmount: workersAmount,
	}, nil
}

type (
	processor struct {
		workersAmount int
	}

	printer struct {
		printDelay time.Duration
	}

	arrayGenerator struct {
		generatorDelay time.Duration
		arrayLength    int
		maxElemValue   int
		minElemValue   int
	}

	Config struct {
		printer
		processor
		arrayGenerator
	}
)

func NewConfig(configPath string) (*Config, error) {
	if configPath == "" {
		return nil, fmt.Errorf("path to config file must be not empty: %w", errors.ErrBadInput)
	}

	if err := godotenv.Load(configPath); err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}

	printer, err := readPrinter()
	if err != nil {
		return nil, fmt.Errorf("error reading printer: %w", err)
	}

	processor, err := readProcessor()
	if err != nil {
		return nil, fmt.Errorf("error reading processor: %w", err)
	}

	arrayGenerator, err := readArrayGenerator()
	if err != nil {
		return nil, fmt.Errorf("error reading arrayGenerator: %w", err)
	}

	return &Config{
		printer:        printer,
		arrayGenerator: arrayGenerator,
		processor:      processor,
	}, nil
}

func (a *printer) PrintDelay() time.Duration {
	return a.printDelay
}

func (a *processor) WorkersAmount() int {
	return a.workersAmount
}

func (a *arrayGenerator) GenerateDelay() time.Duration {
	return a.generatorDelay
}

func (a *arrayGenerator) RandArrLen() int {
	return a.arrayLength
}

func (a *arrayGenerator) RandArrMaxVal() int {
	return a.maxElemValue
}

func (a *arrayGenerator) RandArrMinVal() int {
	return a.minElemValue
}
