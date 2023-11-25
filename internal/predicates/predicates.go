package predicates

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
