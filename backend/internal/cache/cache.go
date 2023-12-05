package cache

import "errors"

type Cache interface {
	Get(id string) (string, error)
	Put(id string, content string) error
	Clear() error
}

var (
	ErrCouldNotRead  = errors.New("could not read from cache")
	ErrCouldNotWrite = errors.New("could not write to cache")
	ErrCouldNotClear = errors.New("could not clear the cache")
)
