package manager

import (
	"testing"

	"github.com/matDobek/gov--attendance-check/internal/testing/assert"
)

func TestStatueParams(t *testing.T) {
  t.Run("builder functions", func(t *testing.T){
    t.Parallel()

    params := NewStatueParams()

    assert.Equal(t, params.VotingNumber.IsNone(), true)
    assert.Equal(t, params.SessionNumber.IsNone(), true)
    assert.Equal(t, params.TermNumber.IsNone(), true)
    assert.Equal(t, params.Title.IsNone(), true)

    params.
      WithTermNumber(1).
      WithSessionNumber(2).
      WithVotingNumber(3).
      WithTitle("foo")

    assert.Equal(t, params.VotingNumber.IsSome(), true)
    assert.Equal(t, params.SessionNumber.IsSome(), true)
    assert.Equal(t, params.TermNumber.IsSome(), true)
    assert.Equal(t, params.Title.IsSome(), true)
  })

  t.Run("validation", func(t *testing.T) {
    t.Parallel()

    params := NewStatueParams()
    valid, errors := params.IsValid()
    errs := errors.(StatueErrors)

    assert.Equal(t, valid, false)
    assert.Equal(t, len(errs.TermNumber), 2)
    assert.Equal(t, len(errs.SessionNumber), 2)
    assert.Equal(t, len(errs.VotingNumber), 2)
  })
}
