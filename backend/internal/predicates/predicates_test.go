package predicates

import (
	"testing"

	"github.com/matDobek/gov--attendance-check/internal/testing/assert"
)

func TestPresent(t *testing.T) {
	t.Run("returns true when value is not empty", func(t *testing.T) {
		t.Parallel()

		tests := [][]any{
			{[]string{}, false},
			{[]string{"a"}, true},
		}

		for _, test := range tests {
			arg0 := test[0].([]string)
			expectation := test[1]
			result := Present(arg0)

			if result != expectation {
				t.Errorf("for %#v, got: %v, expected: %v", arg0, result, expectation)
			}
		}
	})
}

func TestContains(t *testing.T) {
	t.Run("returns true when slice do contain element", func(t *testing.T) {
		t.Parallel()

		got := Contains([]string{"a", "b", "c"}, "b")
		assert.Equal(t, got, true)
	})

	t.Run("returns false when slice does not contain element", func(t *testing.T) {
		t.Parallel()

		got := Contains([]string{"a", "b", "c"}, "z")
		assert.Equal(t, got, false)
	})

	t.Run("returns false when slice and element are empty", func(t *testing.T) {
		t.Parallel()

		got := Contains([]string{}, "")
		assert.Equal(t, got, false)
	})

	t.Run("returns false when slice is empty", func(t *testing.T) {
		t.Parallel()

		got := Contains([]string{}, "a")
		assert.Equal(t, got, false)
	})

	t.Run("returns false when element is empty", func(t *testing.T) {
		t.Parallel()

		got := Contains([]string{"a", "b", "c"}, "")
		assert.Equal(t, got, false)
	})
}

func TestBetween(t *testing.T) {
	t.Run("returns true when number is between two numbers", func(t *testing.T) {
		t.Parallel()

		got := Between(5, 1, 10)
		assert.Equal(t, got, true)
	})

	t.Run("returns true when number is eq the lower bound", func(t *testing.T) {
		t.Parallel()

		got := Between(1, 1, 10)
		assert.Equal(t, got, true)
	})

	t.Run("returns true when number is eq the higher bound", func(t *testing.T) {
		t.Parallel()

		got := Between(10, 1, 10)
		assert.Equal(t, got, true)
	})

	t.Run("returns false when number is outside the bounds", func(t *testing.T) {
		t.Parallel()

		got1 := Between(0, 1, 10)
		got2 := Between(11, 1, 10)

		assert.Equal(t, got1, false)
		assert.Equal(t, got2, false)
	})
}
