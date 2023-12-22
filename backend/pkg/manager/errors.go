package manager

import (
	"errors"
)

var (
	ErrValueRequired  = errors.New("value is required")
	ErrNonZeroValue  = errors.New("expected non empty value")
	ErrPositiveValue = errors.New("expected positive value")
	ErrNegativeValue = errors.New("expected negative zero value")
)
