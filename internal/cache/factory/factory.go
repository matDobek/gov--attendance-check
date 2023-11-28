package factory

import (
	"github.com/matDobek/gov--attendance-check/internal/cache"
	"github.com/matDobek/gov--attendance-check/internal/cache/dummy"
	"github.com/matDobek/gov--attendance-check/internal/cache/files"
)

func DummyCache() (cache.Cache, error) {
	c, err := dummy.New()

	return cache.Cache(&c), err
}

func FileCache() (cache.Cache, error) {
	c, err := files.New()

	return cache.Cache(&c), err
}
