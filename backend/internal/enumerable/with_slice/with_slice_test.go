package withslice

import (
	"testing"
	"time"

	"github.com/matDobek/gov--attendance-check/internal/predicates"
	"github.com/matDobek/gov--attendance-check/internal/testing/assert"
	"github.com/matDobek/gov--attendance-check/internal/testing/logger"
)

func TestNew(t *testing.T) {
	t.Run("it returns struct with proper fields", func(t *testing.T) {
		t.Parallel()

		got, _ := (New([]int{1, 2, 3})).(*SliceEnumerable[int]) // assert type to access underlying fields

		assert.Equal(t, got.ok, []bool{true, true, true})
		assert.Equal(t, got.data, []int{1, 2, 3})
	})
}

func TestWithSlice(t *testing.T) {
	t.Run("it behaves like enumerable", func(t *testing.T) {
		t.Parallel()

		got := New([]int{0, 1, 2}).Map(func(v int) int {
			return v + 1
		}).Map(func(v int) int {
			return v * v
		}).Filter(func(v int) bool {
			return v%2 == 0
		}).Do()

		assert.Equal(t, got, []int{4})
	})

	t.Run("it takes care of each data value in parallel", func(t *testing.T) {
		t.Parallel()

		data := make([]int, 100)

		f := func() {
			_ = New(data).Map(func(v int) int {
				time.Sleep(20 * time.Millisecond)
				return v
			}).Map(func(v int) int {
				time.Sleep(20 * time.Millisecond)
				return v
			}).Do()
		}

		delta := bench(f).Milliseconds()

		if !predicates.Between(delta, 40, 45) {
			logger.LogError(t, "took %v ms", delta)
		}
	})
}

func bench(f func()) time.Duration {
	start := time.Now()
	f()
	delta := time.Since(start)

	return delta
}
