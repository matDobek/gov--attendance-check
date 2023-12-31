package statue_store

import (
	"testing"
	"time"

	"github.com/matDobek/gov--attendance-check/internal/storage"
	"github.com/matDobek/gov--attendance-check/internal/testing/assert"
	"github.com/matDobek/gov--attendance-check/internal/testing/logger"
	"github.com/matDobek/gov--attendance-check/internal/utils"
	"github.com/matDobek/gov--attendance-check/pkg/manager"
)

var (
  dbURL = utils.TestPrimaryDatabaseURL()
)

func TestInsert(t *testing.T) {
  statueStore := NewSQLStore(storage.NewSQLDatabase(dbURL))

  t.Run("saves valid statue", func(t *testing.T) {
    cleanUp(t)

    params := manager.NewStatueParams().
      WithTermNumber(1).
      WithSessionNumber(2).
      WithVotingNumber(3).
      WithTitle("foo")

    time_before := time.Now()
    result, err := statueStore.Insert(*params)
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

func TestAll(t *testing.T) {
  statueStore := NewSQLStore(storage.NewSQLDatabase(dbURL))

  t.Run("returns empty slice when no rows", func(t *testing.T) {
    cleanUp(t)
    result, err := statueStore.All()

    assert.Error(t, err)
    assert.Equal(t, len(result), 0)
  })

  t.Run("returns all existing rows", func(t *testing.T) {
    cleanUp(t)

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

    _, err := statueStore.Insert(*params1)
    assert.Error(t, err)

    _, err = statueStore.Insert(*params2)
    assert.Error(t, err)

    result, err := statueStore.All()

    assert.Error(t, err)
    assert.Equal(t, len(result), 2)
    assert.Equal(t, result[0].Title, "foo")
    assert.Equal(t, result[1].Title, "bar")
  })
}

func cleanUp(t *testing.T) {
  t.Helper()

  t.Cleanup(func() {
    t.Helper()
    db := storage.NewSQLDatabase(dbURL)
    _, err := db.Exec("truncate table statues cascade")
    if err != nil {
      panic(err)
    }
  })
}
