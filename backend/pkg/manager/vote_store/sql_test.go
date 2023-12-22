package vote_store

import (
	"testing"
	"time"

	"github.com/matDobek/gov--attendance-check/internal/storage"
	"github.com/matDobek/gov--attendance-check/internal/testing/assert"
	"github.com/matDobek/gov--attendance-check/internal/testing/logger"
	"github.com/matDobek/gov--attendance-check/internal/utils"
	"github.com/matDobek/gov--attendance-check/pkg/manager"
	"github.com/matDobek/gov--attendance-check/pkg/manager/politician_store"
	"github.com/matDobek/gov--attendance-check/pkg/manager/statue_store"
)

var (
  dbURL = utils.TestPrimaryDatabaseURL()
)

func TestInsert(t *testing.T) {
  db := storage.NewSQLDatabase(dbURL)
  store := NewSQLStore(db)
  p_store := politician_store.NewSQLStore(db)
  s_store := statue_store.NewSQLStore(db)

  t.Run("saves valid", func(t *testing.T) {
    cleanUp(t)

    politician, err := p_store.Insert(*manager.NewPoliticianParams().WithParty("foo").WithName("bar"))
    assert.Error(t, err)

    statue, err := s_store.Insert(*manager.NewStatueParams().
      WithTermNumber(1).
      WithSessionNumber(1).
      WithVotingNumber(1).
      WithTitle("foo"))
    assert.Error(t, err)

    params := manager.NewVoteParams().
      WithResponse("foo").
      WithPoliticianID(politician.ID).
      WithStatueID(statue.ID)

    time_before := time.Now()
    result, err := store.Insert(*params)
    time_after := time.Now()

    assert.Error(t, err)
    assert.Equal(t, result.Response, "foo")
    assert.Equal(t, result.StatueID, statue.ID)
    assert.Equal(t, result.PoliticianID, politician.ID)

    if !time_before.Before(result.CreatedAt) || !time_after.After(result.CreatedAt) {
      logger.LogError(t, "created_at is not valid: %v", result.CreatedAt)
    }

    if !time_before.Before(result.UpdatedAt) || !time_after.After(result.UpdatedAt) {
      logger.LogError(t, "updated_at is not valid: %v", result.UpdatedAt)
    }
  })
}

func TestAll(t *testing.T) {
  db := storage.NewSQLDatabase(dbURL)
  store := NewSQLStore(db)
  p_store := politician_store.NewSQLStore(db)
  s_store := statue_store.NewSQLStore(db)

  t.Run("returns empty slice when no rows", func(t *testing.T) {
    cleanUp(t)
    result, err := store.All()

    assert.Error(t, err)
    assert.Equal(t, len(result), 0)
  })

  t.Run("returns all existing rows", func(t *testing.T) {
    cleanUp(t)

    politician, err := p_store.Insert(*manager.NewPoliticianParams().WithParty("foo").WithName("bar"))
    assert.Error(t, err)

    statue, err := s_store.Insert(*manager.NewStatueParams().
      WithTermNumber(1).
      WithSessionNumber(1).
      WithVotingNumber(1).
      WithTitle("foo"))
    assert.Error(t, err)

    params1 := manager.NewVoteParams().
      WithResponse("foo").
      WithPoliticianID(politician.ID).
      WithStatueID(statue.ID)

    params2 := manager.NewVoteParams().
      WithResponse("bar").
      WithPoliticianID(politician.ID).
      WithStatueID(statue.ID)

    _, err = store.Insert(*params1)
    assert.Error(t, err)

    _, err = store.Insert(*params2)
    assert.Error(t, err)

    result, err := store.All()

    assert.Error(t, err)
    assert.Equal(t, len(result), 2)
    assert.Equal(t, result[0].Response, "foo")
    assert.Equal(t, result[1].Response, "bar")
  })
}

func cleanUp(t *testing.T) {
  t.Helper()

  t.Cleanup(func() {
    t.Helper()
    db := storage.NewSQLDatabase(dbURL)

    _, err := db.Exec("truncate table votes cascade")
    if err != nil {
      panic(err)
    }
    _, err = db.Exec("truncate table politicians cascade")
    if err != nil {
      panic(err)
    }
    _, err = db.Exec("truncate table votes cascade")
    if err != nil {
      panic(err)
    }
  })
}
