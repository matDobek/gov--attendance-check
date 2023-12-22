package manager

import (
	"testing"

	"github.com/matDobek/gov--attendance-check/internal/testing/assert"
)

func TestPoliticianParams(t *testing.T) {
  t.Run("builder functions", func(t *testing.T){
    t.Parallel()

    params := NewPoliticianParams()

    assert.Equal(t, params.Party.IsNone(), true)
    assert.Equal(t, params.Name.IsNone(), true)

    params.WithName("foo").WithParty("bar")

    assert.Equal(t, params.Party.IsSome(), true)
    assert.Equal(t, params.Name.IsSome(), true)
  })

  t.Run("validation", func(t *testing.T) {
    t.Parallel()

    params := NewPoliticianParams()
    valid, errors := params.IsValid()
    errs := errors.(PoliticianErrors)

    assert.Equal(t, valid, false)
    assert.Equal(t, len(errs.Name), 2)
    assert.Equal(t, len(errs.Party), 2)
  })
}
