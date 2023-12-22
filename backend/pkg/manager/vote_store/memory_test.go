package vote_store

import (
	"slices"
	"testing"
	"time"

	"github.com/matDobek/gov--attendance-check/internal/testing/assert"
	"github.com/matDobek/gov--attendance-check/internal/testing/logger"
	"github.com/matDobek/gov--attendance-check/pkg/manager"
)

func TestMemStoreInsert(t *testing.T) {
  t.Run("saves valid Vote", func(t *testing.T) {
    t.Parallel()

    store := NewMemStore()

    params := manager.NewVoteParams().
      WithResponse("foo").
      WithStatueID(1).
      WithPoliticianID(2)

    time_before := time.Now()
    result, err := store.Insert(*params)
    time_after := time.Now()

    assert.Error(t, err)
    assert.Equal(t, result.Response, "foo")
    assert.Equal(t, result.StatueID, 1)
    assert.Equal(t, result.PoliticianID, 2)

    if !time_before.Before(result.CreatedAt) || !time_after.After(result.CreatedAt) {
      logger.LogError(t, "created_at is not valid: %v", result.CreatedAt)
    }

    if !time_before.Before(result.UpdatedAt) || !time_after.After(result.UpdatedAt) {
      logger.LogError(t, "updated_at is not valid: %v", result.UpdatedAt)
    }
  })
}

func TestMemStoreAll(t *testing.T) {
  t.Run("returns empty slice when no rows", func(t *testing.T) {
    t.Parallel()

    store := NewMemStore()

    result, err := store.All()

    assert.Error(t, err)
    assert.Equal(t, len(result), 0)
  })

  t.Run("returns all existing rows", func(t *testing.T) {
    t.Parallel()

    store := NewMemStore()

    params2 := manager.NewVoteParams().
      WithResponse("foo").
      WithStatueID(1).
      WithPoliticianID(1)

    params1 := manager.NewVoteParams().
      WithResponse("bar").
      WithStatueID(1).
      WithPoliticianID(1)

    _, err := store.Insert(*params1)
    assert.Error(t, err)

    _, err = store.Insert(*params2)
    assert.Error(t, err)

    result, err := store.All()

    assert.Error(t, err)
    assert.Equal(t, len(result), 2)

    var responses []string
    for _, r := range result {
      responses = append(responses, r.Response)
    }

    for _, exp := range []string{"foo", "bar"} {
      if !slices.Contains(responses, exp) {
        logger.LogError(t, "%v expected to contain %v", responses, exp)
      }
    }
  })
}
