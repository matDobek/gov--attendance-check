package assert

import (
	"reflect"
	"testing"
)

func Equal(t *testing.T, got, want any) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func Error(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Error(err)
	}
}
