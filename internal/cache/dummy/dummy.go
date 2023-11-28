package dummy

import (
	"github.com/matDobek/gov--attendance-check/internal/cache"
)

//===================
// Types
//===================

type DummyCache struct{}

// verify interface implementation
var _ cache.Cache = (*DummyCache)(nil)

//===================
// Interface Functions
//===================

func New() (DummyCache, error) {
	return DummyCache{}, nil
}

func (_ *DummyCache) Put(id string, content string) error {
	return cache.ErrCouldNotWrite
}

func (_ *DummyCache) Get(id string) (string, error) {
	return "", cache.ErrCouldNotRead
}

func (_ *DummyCache) Clear() error {
	return nil
}
