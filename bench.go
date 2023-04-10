package main

import "github.com/ghetzel/shmtool/shm"

type Bench struct {
	segment *shm.Segment
}

func NewBench() *Bench {
	return &Bench{}
}

func (bench *Bench) Init() {
	bench.segment, _ = shm.Create(1024)
}

func (bench *Bench) Destroy() {
	err := bench.segment.Destroy()
	if err != nil {
		panic(err)
	}
}
