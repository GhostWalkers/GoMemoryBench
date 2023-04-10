package main

import (
	"encoding/binary"
	"math"
)

func (bench *Bench) WriteFloat64Bench() {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, math.Float64bits(0.42))

	bench.segment.Write(b)
}

func (bench *Bench) WriteInt64Bench() {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, 0xDEAD_BEEF)
	bench.segment.Write(b)
}

func (bench *Bench) WriteRawBench() {
	bench.segment.Write([]byte("ABCDEFGH"))
}
