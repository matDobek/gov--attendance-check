package types

import (
	"database/sql/driver"
)

type Maybe[T any] struct {
  value T
  valid bool
}

func None[T any]() Maybe[T] {
  var v T

  return Maybe[T]{ value: v, valid: false }
}

func Some[T any](v T) Maybe[T] {
  return Maybe[T]{ value: v, valid: true }
}

func (m Maybe[T]) IsSome() bool {
  return m.valid
}

func (m Maybe[T]) IsNone() bool {
  return !m.valid
}

func (m Maybe[T]) Unwrap() (T, bool) {
  return m.value, m.valid
}

func (m Maybe[T]) Value() (driver.Value, error) {
  v, ok := m.Unwrap()

  if !ok {
    return nil, nil
  }

  // we must provide `v` that Value is able to handle. In times of writing, those values are:
  //
  // int64
  // float64
  // bool
  // []byte
  // string
  // time.Time
  //
  // ref: https://pkg.go.dev/database/sql/driver#Value

  switch x := any(v).(type) {
  case int:
    return int64(x), nil
  default:
    return v, nil
  }
}
