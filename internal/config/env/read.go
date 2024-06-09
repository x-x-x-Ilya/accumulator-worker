package env

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/x-x-x-Ilya/accumulator-worker/pkg/errors"
)

func mustPositiveDuration(key string) (time.Duration, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return 0, fmt.Errorf("required ENV %q is not set: %w", key, errors.ErrBadInput)
	}

	res, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf("required ENV %q must be time.Duration but it's %q: %w", key, value, err)
	}

	if res <= 0 {
		return 0, fmt.Errorf("required positive duration: %w", errors.ErrBadInput)
	}

	return res, nil
}

func mustPositiveInt(key string) (int, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return 0, fmt.Errorf("required ENV %q is not set: %w", key, errors.ErrBadInput)
	}

	res, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("required ENV %q must be int but it's %q: %w", key, value, err)
	}

	if res <= 0 {
		return 0, fmt.Errorf("required positive integer: %w", errors.ErrBadInput)
	}

	return res, nil
}

func shouldPositiveInt(key string, defaultVal int) (int, error) {
	res, err := shouldInt(key, defaultVal)
	if err != nil {
		return defaultVal, fmt.Errorf("reaing error: %w", err)
	}

	if res <= 0 {
		return 0, fmt.Errorf("required positive integer: %w", errors.ErrBadInput)
	}

	return res, nil
}

func shouldInt(key string, defaultVal int) (int, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultVal, nil
	}

	res, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("required ENV %q must be int but it's %q: %w", key, value, err)
	}

	return res, nil
}
