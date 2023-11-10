

# Struct Case

`LargeStruct v.s. *LargeStruct`

```
go test -bench=. ./structcase/ -benchmem -count=1 -benchtime=100x
```

```
goos: darwin
goarch: arm64
pkg: ref-val/structcase
BenchmarkPassByValue-10              100          14327898 ns/op        268435479 B/op         2 allocs/op
BenchmarkPassByReference-10          100           6789992 ns/op        134217731 B/op         1 allocs/op
PASS
ok      ref-val/structcase      2.583s
```

# Array Case

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

# Convert Array to Another Structure Case


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


# Show Graph

Add this flags ` -cpuprofile=output/cpu.out -memprofile=output/mem.out` to the origin command.
Set up server...
```
go tool pprof -http :8080 output/cpu.out
go tool pprof -http :8081 output/mem.out
```

Examples
@todo: attached images and explain them.