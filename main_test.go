package main

import (
	"testing"
)

func BenchmarkWriteFloat64Bench(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteFloat64Bench()
	}
}

func BenchmarkWriteInt64Bench(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteInt64Bench()
	}
}

func BenchmarkWriteRawBench(b *testing.B) {
	WriteRawBench()
}
