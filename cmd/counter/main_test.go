package main

import (
	"testing"
)

func BenchmarkAtomicCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = AtomicCounter()
	}
}

func BenchmarkUsualCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = UsualCounter()
	}
}
