package main

import (
	"encoding/binary"
	"github.com/ghetzel/shmtool/shm"
	"math"
)

func WriteFloat64Bench() {
	segment, err := shm.Create(1024)
	if err != nil {
		panic(err)
	}
	defer segment.Destroy()

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, math.Float64bits(0.42))
	_, err = segment.Write(b)
	if err != nil {
		panic(err)
	}
}

func WriteInt64Bench() {
	segment, err := shm.Create(1024)
	if err != nil {
		panic(err)
	}
	defer segment.Destroy()

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, 0xDEAD_BEEF)
	_, err = segment.Write(b)
	if err != nil {
		panic(err)
	}
}

func WriteRawBench() {
	segment, err := shm.Create(1024)
	if err != nil {
		panic(err)
	}
	defer segment.Destroy()

	_, err = segment.Write([]byte("ABCDEFGH"))
	if err != nil {
		panic(err)
	}
}
