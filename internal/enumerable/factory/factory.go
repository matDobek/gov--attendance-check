package factory

import (
	"github.com/matDobek/gov--attendance-check/internal/enumerable"
	withslice "github.com/matDobek/gov--attendance-check/internal/enumerable/with_slice"
)

func WithSlice[T any](xs []T) enumerable.Enumerable[T] {
	return withslice.New[T](xs)
}
