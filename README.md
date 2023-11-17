# Go Performance

## Test List

### 001 - Object
Edit person object: `Person v.s. *Person`

```
go test -bench=. ./001/ -benchmem -count=1 -benchtime=100x
```

```
goos: darwin
goarch: arm64
pkg: go-performance/001
BenchmarkPassByValue-10              100                 1.250 ns/op           0 B/op          0 allocs/op
BenchmarkPassByReference-10          100                 2.080 ns/op           0 B/op          0 allocs/op
PASS
ok      go-performance/001      0.100s
```

### 002 - Slice
Edit the slice content: `[]People v.s. []*People`

```
go test -bench=. ./002/ -benchmem -count=1 -benchtime=100x
```

```
goos: darwin
goarch: arm64
pkg: go-performance/002
BenchmarkPassByValue-10              100                51.66 ns/op            0 B/op          0 allocs/op
BenchmarkPassByReference-10          100                42.50 ns/op            0 B/op          0 allocs/op
PASS
ok      go-performance/002      0.346s
```

### 003 - New Slice

Convert people slice to member slice: `[]People -> []Member v.s. []*People -> []*Member`

```
go test -bench=. ./003/ -benchmem -count=1 -benchtime=100x
```

```
goos: darwin
goarch: arm64
pkg: go-performance/003
BenchmarkPassByValue-10              100               142.1 ns/op           896 B/op          1 allocs/op
BenchmarkPassByReference-10          100               774.6 ns/op          1120 B/op         21 allocs/op
PASS
ok      go-performance/003      0.452s
```

### 004 - Map

`map[int]int v.s. map[int]*int`

```
go test -bench=. ./004/ -benchmem -count=1 -benchtime=100x
```

```
goos: darwin
goarch: arm64
pkg: go-performance/004
BenchmarkPassByValue-10              100          11658263 ns/op        641728556 B/op         2 allocs/op
BenchmarkPassByReference-10          100          38284591 ns/op        641728576 B/op         2 allocs/op
PASS
ok      go-performance/004  5.954s
```

### 005 - Clean Architecture + GRPC + Slice + GET

Simulate the query data flow in the clean architecture + GRPC.
Val: `[]Person -> []*ProtoPerson`
Ptr: `[]*Person -> []*ProtoPerson`

```
go test -bench=. ./004/ -benchmem -count=1 -benchtime=100x
```

```
goos: darwin
goarch: arm64
pkg: go-performance/004
BenchmarkValue-10                    100              8915 ns/op            2092 B/op        102 allocs/op
BenchmarkValueAdapter-10             100              2698 ns/op             652 B/op         21 allocs/op
BenchmarkValueUsecase-10             100              5164 ns/op             972 B/op         41 allocs/op
BenchmarkValueHandler-10             100              6230 ns/op            1772 B/op         82 allocs/op
BenchmarkPointer-10                  100              8405 ns/op            2255 B/op        122 allocs/op
BenchmarkPointerAdapter-10           100              3175 ns/op             815 B/op         41 allocs/op
BenchmarkPointerUsecase-10           100              4932 ns/op            1132 B/op         61 allocs/op
BenchmarkPointerHandler-10           100              3391 ns/op            1135 B/op         61 allocs/op
PASS
ok      go-performance/004    0.450s
```

### 006 - Clean Architecture + GRPC + Slice + Create

Simulate the create data flow in the clean architecture + GRPC.
Val: `[]*InputPerson -> []Person -> []*ProtoPerson`
Ptr: `[]*InputPerson -> []*Person -> []*ProtoPerson`

```
go test -bench=. ./005/ -benchmem -count=1 -benchtime=100x
```

```
goos: darwin
goarch: arm64
pkg: go-performance/005
BenchmarkValue-10                    100              3588 ns/op            4440 B/op        102 allocs/op
BenchmarkValueUsecase-10             100              8015 ns/op           34313 B/op         40 allocs/op
BenchmarkValueAdapter-10             100              1473 ns/op            8639 B/op         20 allocs/op
BenchmarkPointer-10                  100              5018 ns/op            6586 B/op        162 allocs/op
BenchmarkPointerUsecase-10           100              6508 ns/op           36243 B/op         80 allocs/op
BenchmarkPointerAdapter-10           100              2250 ns/op            9608 B/op         40 allocs/op
PASS
ok      go-performance/005      0.305s
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