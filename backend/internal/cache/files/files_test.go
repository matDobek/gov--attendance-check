package files

import (
	"os"
	"testing"

	"github.com/matDobek/gov--attendance-check/internal/cache"
	"github.com/matDobek/gov--attendance-check/internal/testing/assert"
)

func TestNew(t *testing.T) {
	t.Run("returns FileCache{...}", func(t *testing.T) {
		t.Parallel()

		path := t.TempDir()
		c, err := ConfigurableNew(path)
		assert.Error(t, err)
		assert.Equal(t, c, FileCache{dir: path})
	})
}

func TestGet(t *testing.T) {
	t.Run("returns value if key do exist", func(t *testing.T) {
		t.Parallel()

		c, err := ConfigurableNew(t.TempDir())
		assert.Error(t, err)

		err = c.Put("a", "aaa")
		assert.Error(t, err)

		err = c.Put("b", "bbb")
		assert.Error(t, err)

		got, err := c.Get("a")
		assert.Error(t, err)
		assert.Equal(t, got, "aaa")
	})

	t.Run("returns ErrCouldNotRead if key does not exist", func(t *testing.T) {
		t.Parallel()

		c, err := ConfigurableNew(t.TempDir())
		assert.Error(t, err)

		_, err = c.Get("non-existing")
		assert.Equal(t, err, cache.ErrCouldNotRead)
	})
}

func TestPut(t *testing.T) {
	t.Run("saves value under key", func(t *testing.T) {
		t.Parallel()

		c, err := ConfigurableNew(t.TempDir())
		assert.Error(t, err)

		err = c.Put("a", "aaa")
		assert.Error(t, err)

		err = c.Put("b", "bbb")
		assert.Error(t, err)

		got, err := c.Get("a")
		assert.Error(t, err)
		assert.Equal(t, got, "aaa")
	})

	t.Run("overwrites value if key already exists", func(t *testing.T) {
		t.Parallel()

		c, err := ConfigurableNew(t.TempDir())
		assert.Error(t, err)

		err = c.Put("a", "aaa")
		assert.Error(t, err)

		err = c.Put("a", "zzz")
		assert.Error(t, err)

		got, err := c.Get("a")
		assert.Error(t, err)
		assert.Equal(t, got, "zzz")
	})
}

func TestClear(t *testing.T) {
	t.Run("overwrites value if key already exists", func(t *testing.T) {
		t.Parallel()

		c, err := ConfigurableNew(t.TempDir())
		assert.Error(t, err)

		tests := [][]any{
			{"a", "1"},
			{"b", "2"},
		}

		for _, test := range tests {
			key := test[0].(string)
			val := test[1].(string)

			err := c.Put(key, val)
			assert.Error(t, err)
		}

		files, err := os.ReadDir(c.dir)
		assert.Error(t, err)
		assert.Equal(t, len(files), 2)

		err = c.Clear()
		assert.Error(t, err)

		files, err = os.ReadDir(c.dir)
		assert.Error(t, err)
		assert.Equal(t, len(files), 0)
	})
}
