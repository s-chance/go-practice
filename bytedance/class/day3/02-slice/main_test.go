package slice

import "testing"

func BenchmarkNoPreAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoPreAlloc(1000)
	}
}

func BenchmarkPreAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PreAlloc(1000)
	}
}
