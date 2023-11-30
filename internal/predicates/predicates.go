package predicates

import "github.com/matDobek/gov--attendance-check/internal/constraints"

func Present[V any](xs []V) bool {
	return len(xs) > 0
}

func Contains[V comparable](xs []V, c V) bool {
	for _, x := range xs {
		if x == c {
			return true
		}
	}

	return false
}

func Between[T constraints.Number](val, a, b T) bool {
	return (a <= val && val <= b)
}
