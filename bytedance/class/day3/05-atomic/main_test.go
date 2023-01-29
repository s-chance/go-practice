package _5_atomic

import "testing"

func BenchmarkAtomicAddOne(b *testing.B) {
	c := &atomicCounter{}
	for i := 0; i < b.N; i++ {
		AtomicAddOne(c)
	}
}

func BenchmarkMutexAddOne(b *testing.B) {
	c := &mutexCounter{}
	for i := 0; i < b.N; i++ {
		MutexAddOne(c)
	}
}
