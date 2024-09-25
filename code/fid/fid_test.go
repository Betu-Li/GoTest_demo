package fid

import (
	"testing"
)

func benchmarkFid(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fid(n)
	}
}

func BenchmarkFid1(b *testing.B) {
	benchmarkFid(b, 1)
}
func BenchmarkFid2(b *testing.B) {
	benchmarkFid(b, 2)
}
func BenchmarkFid3(b *testing.B) {
	benchmarkFid(b, 3)
}
func BenchmarkFid10(b *testing.B) {
	benchmarkFid(b, 10)
}
func BenchmarkFid20(b *testing.B) {
	benchmarkFid(b, 20)
}
func BenchmarkFid40(b *testing.B) {
	benchmarkFid(b, 40)
}
