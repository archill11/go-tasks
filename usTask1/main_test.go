package main

import "testing"

func Benchmark_cCopy(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cCopy(params)
	}
	b.StopTimer()
}
func Benchmark_cString(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cString(params)
	}
	b.StopTimer()
}