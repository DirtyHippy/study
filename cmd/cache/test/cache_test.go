package test

import (
	"github.com/stretchr/testify/assert"
	"study/cmd/cache/helpers"
	"study/cmd/cache/rwmutex"
	"testing"
)

func Test_Cache(t *testing.T) {
	t.Parallel()
	testCache := rwmutex.NewSimpleCache()

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
		helpers.EmulateLoad(t, testCache, parallelFactor)
	})
}
