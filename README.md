
# bench

[![SeriesCI](https://seriesci.com/seriesci/bench/series/master/benchmark.svg)](https://seriesci.com/seriesci/bench/series/master/benchmark)

Go demo application for continuous benchmarks.

## Get benchmark value with unit (ns/op)

```bash
go test -bench=. | grep BenchmarkEfficient | awk '{print $3}'
```

## Live preview

https://seriesci.com/seriesci/demo-benchmark-go
