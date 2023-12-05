package assert

import (
	"errors"
	"reflect"
	"testing"

	"github.com/matDobek/gov--attendance-check/internal/testing/logger"
)

func Equal(t *testing.T, got, want any) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		args := []any{got, want}
		logger.LogError(t, "got %#v, want %#v", args...)
	}
}

func Error(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		logger.LogError(t, "", err.(any))
	}
}

func ErrorIs(t *testing.T, got error, want error) {
	t.Helper()

	if !errors.Is(got, want) {
		args := []any{got, want}
		logger.LogError(t, "got %#v, want %#v", args...)
	}
}
