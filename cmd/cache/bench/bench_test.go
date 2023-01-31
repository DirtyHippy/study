package bench

import (
	"errors"
	"fmt"
	"study/cmd/cache/mutex"
	"study/cmd/cache/rwmutex"
	"study/cmd/cache/storage"
	"sync"
	"testing"
)

const parallelFactor = 10_000_0

func Benchmark_RWMutex_BalancedLoad(b *testing.B) {
	testCache := rwmutex.NewSimpleCache()
	for i := 0; i < b.N; i++ {
		emulateLoad(testCache, parallelFactor)
	}
}

func Benchmark_Mutex_BalancedLoad(b *testing.B) {
	testCache := mutex.NewSimpleCache()
	for i := 0; i < b.N; i++ {
		emulateLoad(testCache, parallelFactor)
	}
}

func Benchmark_RWMutex_IntensiveReadLoad(b *testing.B) {
	testCache := rwmutex.NewSimpleCache()
	for i := 0; i < b.N; i++ {
		emulateReadIntensiveLoad(testCache, parallelFactor)
	}
}

func Benchmark_Mutex_IntensiveReadLoad(b *testing.B) {
	testCache := mutex.NewSimpleCache()
	for i := 0; i < b.N; i++ {
		emulateReadIntensiveLoad(testCache, parallelFactor)
	}
}

func emulateLoad(s storage.Cache, parallelFactor int) {
	wg := sync.WaitGroup{}

	for i := 0; i < parallelFactor; i++ {
		key := fmt.Sprintf("%v-key", i)
		value := fmt.Sprintf("%v-value", i)

		wg.Add(1)
		go func(k, v string) {
			err := s.Set(k, v)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(key, value)

		wg.Add(1)
		go func(k, v string) {
			_, err := s.Get(k)
			if err != nil && !errors.Is(err, storage.ErrNotFound) {
				panic(err)
			}
			wg.Done()
		}(key, value)

		wg.Add(1)
		go func(k, v string) {
			err := s.Delete(k)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(key, value)

		wg.Wait()
	}
}

func emulateReadIntensiveLoad(s storage.Cache, parallelFactor int) {
	wg := sync.WaitGroup{}

	for i := 0; i < parallelFactor/10; i++ {
		key := fmt.Sprintf("%v-key", i)
		value := fmt.Sprintf("%v-value", i)

		wg.Add(1)
		go func(k, v string) {
			err := s.Set(k, v)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(key, value)

		wg.Add(1)
		go func(k, v string) {
			err := s.Delete(k)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(key, value)

		wg.Wait()
	}

	for i := 0; i < parallelFactor; i++ {
		key := fmt.Sprintf("%v-key", i)
		value := fmt.Sprintf("%v-value", i)

		wg.Add(1)
		go func(k, v string) {
			_, err := s.Get(k)
			if err != nil && !errors.Is(err, storage.ErrNotFound) {
				panic(err)
			}
			wg.Done()
		}(key, value)

		wg.Wait()
	}
}
