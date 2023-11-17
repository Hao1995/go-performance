# Map

## 021 - Map

`map[int]int v.s. map[int]*int`

```
go test -bench=. ./map/native/ -benchmem -count=1 -benchtime=100x
```

```
pkg: go-performance/map/native
BenchmarkPassByValue-10              100            568190 ns/op         6417286 B/op          0 allocs/op
BenchmarkPassByReference-10          100            117078 ns/op         6417285 B/op          0 allocs/op
```
