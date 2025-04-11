FastCGo (ARM64) implementation

# About the benchmark

We run benchmarks on

- Server: Ampere Altra Neoverse-N1 @ 2GHz ARM64, Ubuntu 24.04.1
- Go: 1.24.2 linux/arm64

# Run benchmark

`nice -n -20 go test . -bench=. -count=10 | tee bench.txt && benchstat bench.txt`

# Benchmark results

| Test                         | Result      |
|:-----------------------------|:------------|
| CGoEmptyFunction             | 71.08n ± 1% |
| CGoFunction                  | 78.65n ± 1% |
| GoEmptyFunction              | 1.798n ± 1% |
| GoFunction                   | 1.813n ± 1% |
| FastCGoEmptyFunction         | 2.161n ± 1% |
| FastCGoFunctionOnSystemStack | 4.760n ± 1% |