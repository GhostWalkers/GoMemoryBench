# Go SHM Bench

Comparison between:
- `shmtool` (https://github.com/ghetzel/shmtool/shm)

## Installation

- `docker compose up -d`
- `docker compose exec go go test -c -o bench.test github.com/GhostWalkers/GoMemoryBench`
- `docker compose exec go ./bench.test  -test.v -test.paniconexit0 -test.bench . -test.run ^$`

## Results

```sh
goos: linux
goarch: arm64
pkg: github.com/GhostWalkers/GoMemoryBench
BenchmarkWriteFloat64Bench
BenchmarkWriteFloat64Bench-4   	1000000000	         0.0000250 ns/op
BenchmarkWriteInt64Bench
BenchmarkWriteInt64Bench-4     	1000000000	         0.0000174 ns/op
BenchmarkWriteRawBench
BenchmarkWriteRawBench-4       	1000000000	         0.0000166 ns/op
PASS

```