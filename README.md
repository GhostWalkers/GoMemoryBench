# Go SHM Bench

Comparison between:
- `shmtool` (https://github.com/ghetzel/shmtool)

## Installation

- `docker compose up -d`
- `docker compose exec go go test -c -o bench.test github.com/GhostWalkers/GoMemoryBench`
- `docker compose exec go ./bench.test  -test.v -test.paniconexit0 -test.bench . -test.run ^$`

## Results

```sh
goos: linux
goarch: amd64
pkg: github.com/GhostWalkers/GoMemoryBench
cpu: AMD FX(tm)-6300 Six-Core Processor             
BenchmarkWriteFloat64Bench
BenchmarkWriteFloat64Bench-6    26955710                38.68 ns/op
BenchmarkWriteInt64Bench
BenchmarkWriteInt64Bench-6      26795438                40.70 ns/op
BenchmarkWriteRawBench
BenchmarkWriteRawBench-6        26390073                40.81 ns/op
PASS

```