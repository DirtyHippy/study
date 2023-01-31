package bench

import (
	"study/cmd/cache/storage"
	"study/cmd/cache/with_atomic"
	"sync"
	"testing"
)

func Benchmark_MutexWithMetricsAtomic(b *testing.B) {
	c := with_atomic.New()
	for i := 0; i < b.N; i++ {
		emulateLoadWithMetrics(c, parallelFactor)
	}
}

// emulateLoadWithMetrics вспомогательная функция, создает нагрузку на кеш через горутины и проверяет количество записей в кеше
func emulateLoadWithMetrics(cm storage.CacheWithMetrics, parallelFactor int) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		// CacheWithMetrics также реализует интерфейс Cache
		// По этому работает как есть
		emulateLoad(cm, parallelFactor)
		wg.Done()
	}()

	// Добавим забор метрик с кеша
	for i := 0; i < parallelFactor; i++ {
		wg.Add(1)
		go func() {
			_ = cm.TotalAmount()
			wg.Done()
		}()
	}

	wg.Wait()
}
