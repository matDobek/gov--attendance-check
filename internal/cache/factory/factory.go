package factory

import "github.com/matDobek/gov--attendance-check/internal/cache/files"

func FileCache() (files.FileCache, error) {
	return files.New()
}
