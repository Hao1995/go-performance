# Map

## 021 - Map

`map[int]int v.s. map[int]*int`

```
go test -bench=. ./map/021/ -benchmem -count=1 -benchtime=100x
```

```
pkg: go-performance/map/021
BenchmarkPassByValue-10              100          11658263 ns/op        641728556 B/op         2 allocs/op
BenchmarkPassByReference-10          100          38284591 ns/op        641728576 B/op         2 allocs/op
```
