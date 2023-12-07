package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matDobek/gov--attendance-check/internal/db"
	"github.com/matDobek/gov--attendance-check/internal/testing/assert"
)

func TestGetStatues(t *testing.T) {
	statues := []db.Statue{
		{ID: 1, VotingNo: 1, SessionNo: 1, TermNo: 1, Title: "1st voting"},
	}
	politicians := []db.Politician{}
	votes := []db.Vote{}
	store := db.NewGovStore(statues, politicians, votes)
	server := NewGovServer(store)

	t.Run("GET api/v1/statues/", func(t *testing.T) {
		t.Parallel()

		request, _ := http.NewRequest(http.MethodGet, "/api/v1/statues/", nil)
		response := httptest.NewRecorder()

		server.router.ServeHTTP(response, request)

		var got []db.Statue
		err := json.NewDecoder(response.Body).Decode(&got)
		assert.Error(t, err)

		want := []db.Statue{
			{ID: 1, VotingNo: 1, SessionNo: 1, TermNo: 1, Title: "1st voting"},
		}

		assert.Equal(t, response.Header().Get("Content-Type"), "application/json")
		assert.Equal(t, response.Result().Status, "200 OK")
		assert.Equal(t, got, want)
	})
}

func TestGetPoliticians(t *testing.T) {
	statues := []db.Statue{}
	politicians := []db.Politician{
		{ID: 1, Name: "Jan Kowalski", Party: "PiS"},
	}
	votes := []db.Vote{}
	store := db.NewGovStore(statues, politicians, votes)
	server := NewGovServer(store)

	t.Run("GET api/v1/politicians/", func(t *testing.T) {
		t.Parallel()

		request, _ := http.NewRequest(http.MethodGet, "/api/v1/politicians/", nil)
		response := httptest.NewRecorder()

		server.router.ServeHTTP(response, request)

		var got []db.Politician
		err := json.NewDecoder(response.Body).Decode(&got)
		assert.Error(t, err)

		want := []db.Politician{
			{ID: 1, Name: "Jan Kowalski", Party: "PiS"},
		}

		assert.Equal(t, response.Header().Get("Content-Type"), "application/json")
		assert.Equal(t, response.Result().Status, "200 OK")
		assert.Equal(t, got, want)
	})
}

func TestGetVotes(t *testing.T) {
	statues := []db.Statue{}
	politicians := []db.Politician{}
	votes := []db.Vote{
		{ID: 1, PoliticianId: 1, StatueId: 1, Response: db.VoteResponseNo},
	}
	store := db.NewGovStore(statues, politicians, votes)
	server := NewGovServer(store)

	t.Run("GET api/v1/statues/", func(t *testing.T) {
		t.Parallel()

		request, _ := http.NewRequest(http.MethodGet, "/api/v1/votes/", nil)
		response := httptest.NewRecorder()

		server.router.ServeHTTP(response, request)

		var got []db.Vote
		err := json.NewDecoder(response.Body).Decode(&got)
		assert.Error(t, err)

		want := []db.Vote{
			{ID: 1, PoliticianId: 1, StatueId: 1, Response: db.VoteResponseNo},
		}

		assert.Equal(t, response.Header().Get("Content-Type"), "application/json")
		assert.Equal(t, response.Result().Status, "200 OK")
		assert.Equal(t, got, want)
	})
}
