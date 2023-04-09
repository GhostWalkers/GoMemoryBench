package main

import (
	"testing"
)

func BenchmarkWriteFloat64Bench(b *testing.B) {
	WriteFloat64Bench()
}

func BenchmarkWriteInt64Bench(b *testing.B) {
	WriteInt64Bench()
}

func BenchmarkWriteRawBench(b *testing.B) {
	WriteRawBench()
}
