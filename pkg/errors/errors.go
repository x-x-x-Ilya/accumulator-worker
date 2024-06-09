package errors

import (
	"errors"
)

var (
	ErrBadInput            = errors.New("unexpected value")
	ErrApplicationShutdown = errors.New("ErrApplicationShutdown")
	ErrInternalServerError = errors.New("ErrInternalServerError")
)
