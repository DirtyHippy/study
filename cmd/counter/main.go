package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println(UsualCounter())
	fmt.Println(AtomicCounter())
}

func UsualCounter() int {
	wg := sync.WaitGroup{}
	counter := 0
	wg.Add(1)
	go func() {
		for i := 0; i < 10000; i++ {
			counter++
		}
		wg.Done()
	}()
	wg.Wait()
	return counter
}

func AtomicCounter() int32 {
	wg := sync.WaitGroup{}
	counter := int32(0)
	wg.Add(1)
	go func() {
		for i := 0; i < 10000; i++ {
			atomic.AddInt32(&counter, 1)
		}
		wg.Done()
	}()
	wg.Wait()
	return counter
}
