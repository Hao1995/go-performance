# Clean Architecture + GRPC

## GET Slice

Simulate the query data flow in the clean architecture + GRPC.
Val: `[]Person -> []*ProtoPerson`
Ptr: `[]*Person -> []*ProtoPerson`

```
go test -bench=. ./clean_architecture_grpc/get_slice/ -benchmem -count=1 -benchtime=100x
```

```
BenchmarkValue-10                    100              8915 ns/op            2092 B/op        102 allocs/op
BenchmarkValueAdapter-10             100              2698 ns/op             652 B/op         21 allocs/op
BenchmarkValueUsecase-10             100              5164 ns/op             972 B/op         41 allocs/op
BenchmarkValueHandler-10             100              6230 ns/op            1772 B/op         82 allocs/op
BenchmarkPointer-10                  100              8405 ns/op            2255 B/op        122 allocs/op
BenchmarkPointerAdapter-10           100              3175 ns/op             815 B/op         41 allocs/op
BenchmarkPointerUsecase-10           100              4932 ns/op            1132 B/op         61 allocs/op
BenchmarkPointerHandler-10           100              3391 ns/op            1135 B/op         61 allocs/op
```

## Create Slice

Simulate the create data flow in the clean architecture + GRPC.
Val: `[]*InputPerson -> []Person -> []*ProtoPerson`
Ptr: `[]*InputPerson -> []*Person -> []*ProtoPerson`

```
go test -bench=. ./clean_architecture_grpc/create_slice/ -benchmem -count=1 -benchtime=100x
```

```
BenchmarkValue-10                    100              3588 ns/op            4440 B/op        102 allocs/op
BenchmarkValueUsecase-10             100              8015 ns/op           34313 B/op         40 allocs/op
BenchmarkValueAdapter-10             100              1473 ns/op            8639 B/op         20 allocs/op
BenchmarkPointer-10                  100              5018 ns/op            6586 B/op        162 allocs/op
BenchmarkPointerUsecase-10           100              6508 ns/op           36243 B/op         80 allocs/op
BenchmarkPointerAdapter-10           100              2250 ns/op            9608 B/op         40 allocs/op
```