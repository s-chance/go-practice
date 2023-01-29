package _3_map

import "testing"

func BenchmarkNoPreAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoPreAlloc(100)
	}
}

func BenchmarkPreAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PreAlloc(100)
	}
}
