package politician_store

import (
	"testing"
	"time"

	"github.com/matDobek/gov--attendance-check/internal/testing/assert"
	"github.com/matDobek/gov--attendance-check/internal/testing/logger"
	"github.com/matDobek/gov--attendance-check/pkg/manager"
)

func TestMemStoreInsert(t *testing.T) {
  t.Run("saves valid Politician", func(t *testing.T) {
    t.Parallel()

    store := NewMemStore()

    params := manager.NewPoliticianParams().
      WithName("foo").
      WithParty("bar")

    time_before := time.Now()
    result, err := store.Insert(*params)
    time_after := time.Now()

    assert.Error(t, err)
    assert.Equal(t, result.Name, "foo")
    assert.Equal(t, result.Party, "bar")

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

    params1 := manager.NewPoliticianParams().
      WithName("foo").
      WithParty("foo")

    params2 := manager.NewPoliticianParams().
      WithName("bar").
      WithParty("bar")

    _, err := store.Insert(*params1)
    assert.Error(t, err)

    _, err = store.Insert(*params2)
    assert.Error(t, err)

    result, err := store.All()

    assert.Error(t, err)
    assert.Equal(t, len(result), 2)
    assert.Equal(t, result[0].Name, "foo")
    assert.Equal(t, result[1].Name, "bar")
  })
}
