package statue_store

import (
	"testing"

	"github.com/matDobek/gov--attendance-check/internal/storage"
	"github.com/matDobek/gov--attendance-check/internal/testing/assert"
	"github.com/matDobek/gov--attendance-check/internal/utils"
	"github.com/matDobek/gov--attendance-check/pkg/manager"
)

var (
  dbURL = utils.TestPrimaryDatabaseURL()
)

func TestInsert(t *testing.T) {
  statueStore := NewSQLStatueStore(storage.NewStorage(dbURL))

  t.Run("saves valid statue", func(t *testing.T) {
    cleanUp(t)
    result, err := statueStore.Insert(manager.Statue{
      TermNumber: 1,
      SessionNumber: 2,
      VotingNumber: 3,
      Title: "foo",
    })

    assert.Error(t, err)
    assert.Equal(t, result.TermNumber, 1)
    assert.Equal(t, result.SessionNumber, 2)
    assert.Equal(t, result.VotingNumber, 3)
    assert.Equal(t, result.Title, "foo")
  })
}

func TestALl(t *testing.T) {
  statueStore := NewSQLStatueStore(storage.NewStorage(dbURL))

  t.Run("returns empty slice when no rows", func(t *testing.T) {
    cleanUp(t)
    result, err := statueStore.All()

    assert.Error(t, err)
    assert.Equal(t, len(result), 0)
  })

  t.Run("returns all existing rows", func(t *testing.T) {
    cleanUp(t)
    _, err := statueStore.Insert(manager.Statue{
      TermNumber: 1,
      SessionNumber: 1,
      VotingNumber: 1,
      Title: "foo",
    })
    assert.Error(t, err)

    _, err = statueStore.Insert(manager.Statue{
      TermNumber: 1,
      SessionNumber: 1,
      VotingNumber: 2,
      Title: "bar",
    })
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
    s := storage.NewStorage(dbURL)
    _, err := s.PrimaryDB.Exec("truncate table statues cascade")
    if err != nil {
      panic(err)
    }
  })
}
