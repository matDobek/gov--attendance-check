package types

import (
	"testing"

	"github.com/matDobek/gov--attendance-check/internal/testing/assert"
)

func TestMaybe(t *testing.T){
  t.Run("None values", func (t *testing.T) {
    t.Parallel()

    m := None[int]()

    assert.Equal(t, m.IsNone(), true)
    assert.Equal(t, m.IsSome(), false)

    v, b := m.Unwrap()
    assert.Equal(t, v, 0)
    assert.Equal(t, b, false)
  })

  t.Run("Some with basic type", func (t *testing.T) {
    t.Parallel()

    m := Some(1)

    assert.Equal(t, m.IsNone(), false)
    assert.Equal(t, m.IsSome(), true)

    v, b := m.Unwrap()
    assert.Equal(t, v, 1)
    assert.Equal(t, b, true)
  })

  t.Run("Some with complex type", func (t *testing.T) {
    t.Parallel()

    type user struct {
      firstName string
      lastName string
    }

    val := user{"foo", "bar"}

    m := Some(val)

    assert.Equal(t, m.IsNone(), false)
    assert.Equal(t, m.IsSome(), true)

    v, b := m.Unwrap()
    assert.Equal(t, v, val)
    assert.Equal(t, b, true)
  })

  t.Run("Unwrap returns value, not reference", func (t *testing.T) {
    t.Parallel()

    type user struct {
      firstName string
      lastName string
    }

    m := Some(user{"foo", "bar"})
    v1, _ := m.Unwrap()

    v1.firstName = "baz"
    v2, _ := m.Unwrap()
    assert.Equal(t, v2.firstName, "foo")
  })
}
