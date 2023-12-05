package withslice

import "github.com/matDobek/gov--attendance-check/internal/enumerable"

type SliceEnumerable[T any] struct {
	data []T
	ok   []bool
	fs   [](func(T, bool) (T, bool))
}

var _ enumerable.Enumerable[int] = (*SliceEnumerable[int])(nil)

func New[T any](xs []T) enumerable.Enumerable[T] {
	data := make([]T, len(xs))
	copy(data, xs)

	ok := make([]bool, len(xs))
	for i := range ok {
		ok[i] = true
	}

	return &SliceEnumerable[T]{
		data: data,
		ok:   ok,
		fs:   [](func(T, bool) (T, bool)){},
	}
}

func (e *SliceEnumerable[T]) Map(f func(T) T) enumerable.Enumerable[T] {
	e.fs = append(e.fs, func(val T, ok bool) (T, bool) {
		return f(val), ok
	})

	return e
}

func (e *SliceEnumerable[T]) Filter(f func(T) bool) enumerable.Enumerable[T] {
	e.fs = append(e.fs, func(val T, ok bool) (T, bool) {
		return val, (ok && f(val))
	})

	return e
}

func (e *SliceEnumerable[T]) Do() []T {
	ch := make(chan struct{}, len(e.data))

	for i := 0; i < len(e.data); i++ {
		index := i

		go func() {
			for j := 0; j < len(e.fs); j++ {
				oldV := e.data[index]
				oldOk := e.ok[index]

				if oldOk {
					v, ok := e.fs[j](oldV, oldOk)
					e.data[index] = v
					e.ok[index] = ok
				}
			}

			ch <- struct{}{}
		}()
	}

	for i := 0; i < len(e.data); i++ {
		<-ch
	}

	out := []T{}
	for i := 0; i < len(e.data); i++ {
		if e.ok[i] {
			out = append(out, e.data[i])
		}
	}

	return out
}
