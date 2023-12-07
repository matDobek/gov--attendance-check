package db

import (
	"testing"

	"github.com/matDobek/gov--attendance-check/internal/testing/assert"
)

var (
	statues = []Statue{
		{1, 1, 1, 1, "1st voting"},
		{2, 1, 1, 1, "2nd voting"},
	}
	politicians = []Politician{
		{11, "Jan Kowalski", "PiS"},
		{12, "Adam Nowak", "KO"},
	}
	votes = []Vote{
		{ID: 111, StatueId: 1, PoliticianId: 11, Response: VoteResponseYes},
		{ID: 112, StatueId: 1, PoliticianId: 12, Response: VoteResponseNo},
	}
)

func initDB() *GovStore {
	return NewGovStore(statues, politicians, votes)
}

func TestGetStatues(t *testing.T) {
	d := initDB()

	t.Run("returns all Statues", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, d.GetStatues(), statues)
	})
}

func TestGetVotes(t *testing.T) {
	d := initDB()

	t.Run("returns all Votes", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, d.GetVotes(), votes)
	})
}

func TestGetPoliticians(t *testing.T) {
	d := initDB()

	t.Run("returns all Politicians", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, d.GetPoliticians(), politicians)
	})
}
