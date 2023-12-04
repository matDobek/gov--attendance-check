package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matDobek/gov--attendance-check/internal/db"
	"github.com/matDobek/gov--attendance-check/internal/testing/assert"
)

func TestHandler(t *testing.T) {
	statues := []db.Statue{
		{1, 1, 1, "1st voting", []db.Vote{}},
	}
	store := db.NewGovStore(statues)
	server := NewGovServer(store)

	t.Run("GET foo score", func(t *testing.T) {
		t.Parallel()

		request, _ := http.NewRequest(http.MethodGet, "/statues/", nil)
		response := httptest.NewRecorder()

		server.router.ServeHTTP(response, request)

		var got []db.Statue
		err := json.NewDecoder(response.Body).Decode(&got)
		assert.Error(t, err)

		want := []db.Statue{
			{1, 1, 1, "1st voting", []db.Vote{}},
		}
		assert.Equal(t, got, want)
	})
}
