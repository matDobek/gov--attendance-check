package cache

import "errors"

var (
	ErrCouldNotRead  = errors.New("could not read from cache")
	ErrCouldNotWrite = errors.New("could not write to cache")
	ErrCouldNotClear = errors.New("could not clear the cache")
)
