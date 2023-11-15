# Go Performance

## Various Type Test
### Struct Case

`Person v.s. *Person`

```
go test -bench=. ./structcase/ -benchmem -count=1 -benchtime=100x
```

```
goos: darwin
goarch: arm64
pkg: go-performance/structcase
BenchmarkPassByValue-10              100                 1.670 ns/op           0 B/op          0 allocs/op
BenchmarkPassByReference-10          100                 2.500 ns/op           0 B/op          0 allocs/op
PASS
ok      go-performance/structcase       0.316s
```

### Array Case

`[]int v.s. []*int`

```
go test -bench=. ./refarr/ -benchmem -count=1 -benchtime=100x
```

```
goos: darwin
goarch: arm64
pkg: ref-val/refarr
BenchmarkPassByValue-10              100           7452440 ns/op        134217925 B/op         1 allocs/op
BenchmarkPassByReference-10          100         142818432 ns/op        268436020 B/op  16777217 allocs/op
PASS
ok      ref-val/refarr  15.520s
```

### Convert Array to Another Structure Case

`[]int -> []Tmp v.s. []*int -> []*Tmp`

```
go test -bench=. ./refarrconv/ -benchmem -count=1 -benchtime=100x
```

```
goos: darwin
goarch: arm64
pkg: ref-val/refarrconv
BenchmarkPassByValue-10              100           9628510 ns/op        268435656 B/op         2 allocs/op
BenchmarkPassByReference-10          100         153143861 ns/op        402653785 B/op  16777218 allocs/op
PASS
ok      ref-val/refarrconv      16.730s
```

### Convert Array to Another Structure by Private Func returning Pointer Case

`func xxx(i int) *Tmp`
`[]int -> []Tmp v.s. []*int -> []*Tmp`

```
go test -bench=. ./convbyfuncptr/ -benchmem -count=1 -benchtime=100x
```

```
goos: darwin
goarch: arm64
pkg: go-performance/convbyfuncptr
BenchmarkPassByValue-10              100                39.59 ns/op            8 B/op          1 allocs/op
BenchmarkPassByReference-10          100                66.25 ns/op           16 B/op          2 allocs/op
PASS
ok      go-performance/convbyfuncptr    0.377s
```

### Convert Slice to Slice by Pointer Case

`[]int -> []*Tmp v.s. []*int -> []*Tmp`

```
go test -bench=. ./convtoarrptr/ -benchmem -count=1 -benchtime=100x
```

```
goos: darwin
goarch: arm64
pkg: go-performance/convtoarrptr
BenchmarkPassByValue-10              100                64.17 ns/op           16 B/op          2 allocs/op
BenchmarkPassByReference-10          100                82.08 ns/op           16 B/op          2 allocs/op
PASS
ok      go-performance/convtoarrptr     0.415s
```

### Slice by Value v.s. Slice by Pointer

`[]Tmp v.s. []*Tmp`

```
go test -bench=. ./arrptrmutitime/ -benchmem -count=1 -benchtime=100x
```

```
goos: darwin
goarch: arm64
pkg: go-performance/arrptrmutitime
BenchmarkValue-10            100           1261520 ns/op          314519 B/op      39488 allocs/op
BenchmarkPointer-10          100           1110082 ns/op          314439 B/op      39488 allocs/op
PASS
ok      go-performance/arrptrmutitime   0.546s
```

### Clean Architecture Case

`[]Person v.s. []*Person`

```
go test -bench=. ./cleanarchcase/ -benchmem -count=1 -benchtime=100x
```

```
goos: darwin
goarch: arm64
pkg: go-performance/cleanarchcase
BenchmarkValue-10            100              8683 ns/op            2092 B/op        102 allocs/op
BenchmarkPointer-10          100              9363 ns/op            2252 B/op        122 allocs/op
PASS
ok      go-performance/cleanarchcase    0.277s
```

### Clean Architecture With Edit Input Case

`[]Person v.s. []*Person`

```
go test -bench=. ./cleanarcheditsrccase/ -benchmem -count=1 -benchtime=100x
```

```
goos: darwin
goarch: arm64
pkg: go-performance/cleanarcheditsrccase
BenchmarkValue-10            100              8902 ns/op            1772 B/op        101 allocs/op
BenchmarkPointer-10          100              8792 ns/op            1775 B/op        101 allocs/op
PASS
ok      go-performance/cleanarcheditsrccase     0.411s
```

### Map Case

`map[int]int v.s. map[int]*int`

```
go test -bench=. ./mapcase/ -benchmem -count=1 -benchtime=100x
```

```
goos: darwin
goarch: arm64
pkg: go-performance/mapcase
BenchmarkPassByValue-10              100          11658263 ns/op        641728556 B/op         2 allocs/op
BenchmarkPassByReference-10          100          38284591 ns/op        641728576 B/op         2 allocs/op
PASS
ok      go-performance/mapcase  5.954s
```

# Show Graph

Add this flags ` -cpuprofile=output/cpu.out -memprofile=output/mem.out` to the origin command.
Set up server...
```
go tool pprof -http :8080 output/cpu.out
go tool pprof -http :8081 output/mem.out
```

Enter pprof mode
```
go tool pprof output/cpu.out
(pprof) web
(pprof) text
```

Examples
@todo: attached images and explain them.