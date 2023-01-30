package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println(MutexCounter())
	fmt.Println(AtomicCounter())
}

func MutexCounter() int {
	goroutinesCount := 0
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			goroutinesCount++
			mu.Unlock()
			time.Sleep(time.Microsecond)
			mu.Lock()
			goroutinesCount--
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	return goroutinesCount
}

func AtomicCounter() int32 {
	goroutinesCount := int32(0)
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&goroutinesCount, 1)
			time.Sleep(time.Microsecond)
			atomic.AddInt32(&goroutinesCount, -1)
			wg.Done()
		}()
	}

	wg.Wait()
	return goroutinesCount
}
