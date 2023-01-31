package helpers

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"study/cmd/cache/storage"
	"sync"
	"testing"
)

func EmulateLoad(t *testing.T, s storage.Cache, parallelFactor int) {
	wg := sync.WaitGroup{}

	for i := 0; i < parallelFactor; i++ {
		key := fmt.Sprintf("%v-key", i)
		value := fmt.Sprintf("%v-value", i)

		wg.Add(1)
		go func(k, v string) {
			err := s.Set(k, v)
			assert.NoError(t, err)
			wg.Done()
		}(key, value)

		wg.Add(1)
		go func(k, v string) {
			storedValue, err := s.Get(k)
			if !errors.Is(err, storage.ErrNotFound) {
				assert.Equal(t, storedValue, v)
			}
			wg.Done()
		}(key, value)

		wg.Add(1)
		go func(k, v string) {
			err := s.Delete(k)
			assert.NoError(t, err)
			wg.Done()
		}(key, value)

		wg.Wait()
	}
}

// emulateLoadWithMetrics вспомогательная функция, создает нагрузку на кеш через горутины и проверяет количество записей в кеше
func EmulateLoadWithMetrics(t *testing.T, cm storage.CacheWithMetrics, parallelFactor int) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		// CacheWithMetrics также реализует интерфейс Cache
		// По этому работает как есть
		EmulateLoad(t, cm, parallelFactor)
		wg.Done()
	}()

	// Добавим забор метрик с кеша
	var min, max int64
	for i := 0; i < parallelFactor; i++ {
		wg.Add(1)
		go func() {
			total := cm.TotalAmount()
			if total > max {
				max = total
			}
			if total < min {
				min = total
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(max, min)
}
