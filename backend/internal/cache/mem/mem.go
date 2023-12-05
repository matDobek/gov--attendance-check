package mem

import (
	"github.com/matDobek/gov--attendance-check/internal/cache"
)

//===================
// Types
//===================

type MemCache struct {
	cache map[string]string
}

var _ cache.Cache = (*MemCache)(nil)

//===================
// Functions
//===================

func New() (MemCache, error) {
	return MemCache{make(map[string]string)}, nil
}

//===================
// Interface Functions
//===================

func (c *MemCache) Put(id string, val string) error {
	c.cache[id] = val

	return nil
}

func (c *MemCache) Get(id string) (string, error) {
	val, ok := c.cache[id]

	if !ok {
		return "", cache.ErrCouldNotRead
	}

	return val, nil
}

func (c *MemCache) Clear() error {
	c.cache = make(map[string]string)

	return nil
}
