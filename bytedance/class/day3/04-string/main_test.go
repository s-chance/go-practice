package _4_string

import "testing"

func BenchmarkPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Plus(100, "a")
	}
}

func BenchmarkStrBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StrBuilder(100, "a")
	}
}

func BenchmarkBytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BytesBuffer(100, "a")
	}
}

func BenchmarkPreStrBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PreStrBuilder(100, "a")
	}
}

func BenchmarkPreBytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PreBytesBuffer(100, "a")
	}
}
