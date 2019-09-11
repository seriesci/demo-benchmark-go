
# bench

Go demo application for continuous benchmarks.

## Get benchmark value with unit (ns/op)

```bash
go test -bench=. | grep BenchmarkEfficient | awk '{print $3}'
```