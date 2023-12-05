package enumerable

type Enumerable[T any] interface {
	Map(func(T) T) Enumerable[T]
	Filter(func(T) bool) Enumerable[T]
	Do() []T
}
