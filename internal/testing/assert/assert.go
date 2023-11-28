package assert

import (
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
