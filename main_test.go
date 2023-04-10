package main

import (
	"testing"
)

func BenchmarkWriteFloat64Bench(b *testing.B) {
	bench := NewBench()
	bench.Init()
	defer bench.Destroy()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bench.WriteFloat64Bench()
	}
}

func BenchmarkWriteInt64Bench(b *testing.B) {
	bench := NewBench()
	bench.Init()
	defer bench.Destroy()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bench.WriteInt64Bench()
	}
}

func BenchmarkWriteRawBench(b *testing.B) {
	bench := NewBench()
	bench.Init()
	defer bench.Destroy()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bench.WriteRawBench()
	}
}
