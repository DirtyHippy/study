package main

import "testing"

func BenchmarkAtomicCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = AtomicCounter()
	}
}

func BenchmarkMutexCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MutexCounter()
	}
}
