package files

import (
	"os"
	"strings"
	"testing"

	"github.com/matDobek/gov--attendance-check/internal/testing/assert"
)

var (
	defaultTestPath = []string{"tmp", "test", "cache"}
)

func TestNew(t *testing.T) {
	c, err := New(defaultTestPath)
	assert.Error(t, err)

	projectRoot, err := lookupMod()
	assert.Error(t, err)
	fullPath := strings.Join(append(projectRoot, defaultTestPath...), "/")

	// it creates proper direcotry structure
	{
		_, err := os.Stat(fullPath)
		assert.Error(t, err)
	}

	// it returns proper result
	assert.Equal(t, c, FileCache{dir: fullPath})
}

func TestPutGet(t *testing.T) {
	c, err := New(defaultTestPath)
	assert.Error(t, err)

	tests := [][]any{
		{"a", "1"},
		{"b", "2"},
		{"a", "3"}, // allow overwrites
	}

	for _, test := range tests {
		key := test[0].(string)
		val := test[1].(string)

		err = c.Put(key, val)
		assert.Error(t, err)

		rVal, err := c.Get(key)
		assert.Error(t, err)
		assert.Equal(t, rVal, val)
	}

	err = c.Clear()
	assert.Error(t, err)
}

func TestClear(t *testing.T) {
	c, err := New(defaultTestPath)
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
}
