package statue_store

import (
	"testing"
	"time"

	"github.com/matDobek/gov--attendance-check/internal/testing/assert"
	"github.com/matDobek/gov--attendance-check/internal/testing/logger"
	"github.com/matDobek/gov--attendance-check/pkg/manager"
)

func TestMemStoreInsert(t *testing.T) {
  t.Run("saves valid statue", func(t *testing.T) {
    t.Parallel()

    store := NewMemStore()

    params := manager.NewStatueParams().
      WithTermNumber(1).
      WithSessionNumber(2).
      WithVotingNumber(3).
      WithTitle("foo")

    time_before := time.Now()
    result, err := store.Insert(*params)
    time_after := time.Now()

    assert.Error(t, err)
    assert.Equal(t, result.TermNumber, 1)
    assert.Equal(t, result.SessionNumber, 2)
    assert.Equal(t, result.VotingNumber, 3)
    assert.Equal(t, result.Title, "foo")

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

    params1 := manager.NewStatueParams().
      WithTermNumber(1).
      WithSessionNumber(1).
      WithVotingNumber(1).
      WithTitle("foo")

    params2 := manager.NewStatueParams().
      WithTermNumber(1).
      WithSessionNumber(1).
      WithVotingNumber(2).
      WithTitle("bar")

    _, err := store.Insert(*params1)
    assert.Error(t, err)

    _, err = store.Insert(*params2)
    assert.Error(t, err)

    result, err := store.All()

    assert.Error(t, err)
    assert.Equal(t, len(result), 2)
    assert.Equal(t, result[0].Title, "foo")
    assert.Equal(t, result[1].Title, "bar")
  })
}
