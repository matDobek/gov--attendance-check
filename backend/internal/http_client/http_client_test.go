package httpclient

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matDobek/gov--attendance-check/internal/cache/factory"
	"github.com/matDobek/gov--attendance-check/internal/testing/assert"
)

func TestGet(t *testing.T) {
	t.Run("return response body when server responds", func(t *testing.T) {
		t.Parallel()

		handler := func(w http.ResponseWriter, r *http.Request) {
      _, err := w.Write([]byte("foo"))
      assert.Error(t, err)
		}

		server := httptest.NewServer(http.HandlerFunc(handler))
		defer server.Close()

		cache, err := factory.DummyCache()
		assert.Error(t, err)

		client := New(cache)
		resp, err := client.Get(server.URL)
		assert.Error(t, err)

		assert.Equal(t, string(resp), "foo")
	})

	t.Run("return error when cannot connect to the server", func(t *testing.T) {
		t.Parallel()

		cache, err := factory.DummyCache()
		assert.Error(t, err)

		client := New(cache)
		_, err = client.Get("http://localhost:1234")
		assert.ErrorIs(t, err, ConnectionError{})
	})
}

func TestCachedGet(t *testing.T) {
	t.Run("return cached response if possible", func(t *testing.T) {
		t.Parallel()

		var i int

		handler := func(w http.ResponseWriter, r *http.Request) {
			i++

      _, err := w.Write([]byte("foo"))
      assert.Error(t, err)
		}

		server := httptest.NewServer(http.HandlerFunc(handler))
		defer server.Close()

		cache, err := factory.MemCache()
		assert.Error(t, err)

		client := New(cache)

		resp1, err := client.CachedGet(server.URL)
		assert.Error(t, err)

		resp2, err := client.CachedGet(server.URL)
		assert.Error(t, err)

		assert.Equal(t, string(resp1), "foo")
		assert.Equal(t, string(resp2), "foo")
		assert.Equal(t, i, 1)
	})
}
