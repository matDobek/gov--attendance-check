package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matDobek/gov--attendance-check/internal/testing/assert"
	"github.com/matDobek/gov--attendance-check/pkg/manager"
	"github.com/matDobek/gov--attendance-check/pkg/manager/politician_store"
	"github.com/matDobek/gov--attendance-check/pkg/manager/statue_store"
	"github.com/matDobek/gov--attendance-check/pkg/manager/vote_store"
)

func TestGetStatues(t *testing.T) {
	t.Run("GET api/v1/statues/", func(t *testing.T) {
		t.Parallel()

		// prep data

		statueStore := statue_store.NewMemStore()
		politicianStore := politician_store.NewMemStore()
		voteStore := vote_store.NewMemStore()

		server := NewGovServer(statueStore, politicianStore, voteStore)

		params1 := manager.NewStatueParams().
			WithTitle("foo").
			WithTermNumber(1).
			WithSessionNumber(1).
			WithVotingNumber(1)

		params2 := manager.NewStatueParams().
			WithTitle("bar").
			WithTermNumber(2).
			WithSessionNumber(2).
			WithVotingNumber(2)

		_, err := statueStore.Insert(*params1)
		assert.Error(t, err)

		_, err = statueStore.Insert(*params2)
		assert.Error(t, err)

		// test

		request, _ := http.NewRequest(http.MethodGet, "/api/v1/statues/", nil)
		response := httptest.NewRecorder()

		server.router.ServeHTTP(response, request)

		var got []manager.Statue
		err = json.NewDecoder(response.Body).Decode(&got)
		assert.Error(t, err)

		// validate results

		assert.Equal(t, response.Header().Get("Content-Type"), "application/json")
		assert.Equal(t, response.Result().Status, "200 OK")

		for _, g := range got {
			var included bool

			for _, title := range []string{"foo", "bar"} {
				if title == g.Title {
					included = true
				}
			}

			assert.Equal(t, included, true)
		}
	})
}

func TestGetPoliticians(t *testing.T) {
	t.Run("GET api/v1/politicians/", func(t *testing.T) {
		t.Parallel()

		// prep data

		statueStore := statue_store.NewMemStore()
		politicianStore := politician_store.NewMemStore()
		voteStore := vote_store.NewMemStore()

		server := NewGovServer(statueStore, politicianStore, voteStore)

		params1 := manager.NewPoliticianParams().WithName("foo").WithParty("foo")
		params2 := manager.NewPoliticianParams().WithName("bar").WithParty("bar")

		_, err := politicianStore.Insert(*params1)
		assert.Error(t, err)

		_, err = politicianStore.Insert(*params2)
		assert.Error(t, err)

		// test

		request, _ := http.NewRequest(http.MethodGet, "/api/v1/politicians/", nil)
		response := httptest.NewRecorder()

		server.router.ServeHTTP(response, request)

		var got []manager.Politician
		err = json.NewDecoder(response.Body).Decode(&got)
		assert.Error(t, err)

		// validate results

		assert.Equal(t, response.Header().Get("Content-Type"), "application/json")
		assert.Equal(t, response.Result().Status, "200 OK")

		for _, g := range got {
			var included bool

			for _, v := range []string{"foo", "bar"} {
				if v == g.Name {
					included = true
				}
			}

			assert.Equal(t, included, true)
		}
	})
}

func TestGetVotes(t *testing.T) {
	t.Run("GET api/v1/votes/", func(t *testing.T) {
		t.Parallel()

		// prep data

		statueStore := statue_store.NewMemStore()
		politicianStore := politician_store.NewMemStore()
		voteStore := vote_store.NewMemStore()

		server := NewGovServer(statueStore, politicianStore, voteStore)

		params1 := manager.NewVoteParams().WithResponse("foo").WithStatueID(1).WithPoliticianID(1)
		params2 := manager.NewVoteParams().WithResponse("bar").WithStatueID(2).WithPoliticianID(2)

		_, err := voteStore.Insert(*params1)
		assert.Error(t, err)

		_, err = voteStore.Insert(*params2)
		assert.Error(t, err)

		// test

		request, _ := http.NewRequest(http.MethodGet, "/api/v1/votes/", nil)
		response := httptest.NewRecorder()

		server.router.ServeHTTP(response, request)

		var got []manager.Vote
		err = json.NewDecoder(response.Body).Decode(&got)
		assert.Error(t, err)

		// validate results

		assert.Equal(t, response.Header().Get("Content-Type"), "application/json")
		assert.Equal(t, response.Result().Status, "200 OK")

		for _, g := range got {
			var included bool

			for _, v := range []string{"foo", "bar"} {
				if v == g.Response {
					included = true
				}
			}

			assert.Equal(t, included, true)
		}
	})
}
