package manager

import (
	"testing"

	"github.com/matDobek/gov--attendance-check/internal/testing/assert"
)

func TestVoteParams(t *testing.T) {
  t.Run("builder functions", func(t *testing.T){
    t.Parallel()

    params := NewVoteParams()

    assert.Equal(t, params.PoliticianID.IsNone(), true)
    assert.Equal(t, params.StatueID.IsNone(), true)
    assert.Equal(t, params.Response.IsNone(), true)

    params.
      WithStatueID(2).
      WithPoliticianID(3).
      WithResponse("foo")

    assert.Equal(t, params.PoliticianID.IsSome(), true)
    assert.Equal(t, params.StatueID.IsSome(), true)
    assert.Equal(t, params.Response.IsSome(), true)
  })

  t.Run("validation", func(t *testing.T) {
    t.Parallel()

    params := NewVoteParams()
    valid, errors := params.IsValid()
    errs := errors.(VoteErrors)

    assert.Equal(t, valid, false)
    assert.Equal(t, len(errs.StatueID), 2)
    assert.Equal(t, len(errs.PoliticianID), 2)
  })
}
