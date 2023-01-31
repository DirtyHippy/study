package test

import (
	"github.com/stretchr/testify/assert"
	"study/cmd/cache/helpers"
	"study/cmd/cache/with_atomic"
	"testing"
)

func Test_CacheWithMetrics(t *testing.T) {
	t.Parallel()
	// Разные имплементации кешей
	testCache := with_atomic.New()
	//testCache := no_mutex.New()

	t.Run("correctly stored value", func(t *testing.T) {
		t.Parallel()
		key := "someKey"
		value := "someValue"

		err := testCache.Set(key, value)
		assert.NoError(t, err)
		storedValue, err := testCache.Get(key)
		assert.NoError(t, err)

		assert.Equal(t, value, storedValue)
	})

	t.Run("no data races", func(t *testing.T) {
		t.Parallel()

		parallelFactor := 100_000
		helpers.EmulateLoadWithMetrics(t, testCache, parallelFactor)
	})
}
