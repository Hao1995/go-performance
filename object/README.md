
# Object

## Native
Edit person object: `Person v.s. *Person`

```
go test -bench=. ./object/native/ -benchmem -count=5 -benchtime=100x
```

```
BenchmarkPassByValue-10              100                 1.670 ns/op           0 B/op          0 allocs/op
BenchmarkPassByValue-10              100                 1.670 ns/op           0 B/op          0 allocs/op
BenchmarkPassByValue-10              100                 0.8400 ns/op          0 B/op          0 allocs/op
BenchmarkPassByValue-10              100                 1.250 ns/op           0 B/op          0 allocs/op
BenchmarkPassByValue-10              100                 1.670 ns/op           0 B/op          0 allocs/op
BenchmarkPassByReference-10          100                 1.250 ns/op           0 B/op          0 allocs/op
BenchmarkPassByReference-10          100                 1.250 ns/op           0 B/op          0 allocs/op
BenchmarkPassByReference-10          100                 0.8300 ns/op          0 B/op          0 allocs/op
BenchmarkPassByReference-10          100                 0.8400 ns/op          0 B/op          0 allocs/op
BenchmarkPassByReference-10          100                 0.8300 ns/op          0 B/op          0 allocs/op
```

## New Object
Return new object

```
go test -bench=. ./object/new_object/ -benchmem -count=1 -benchtime=100x
```

```
BenchmarkPassByValue-10              100                 1.670 ns/op           0 B/op          0 allocs/op
BenchmarkPassByValue-10              100                 1.250 ns/op           0 B/op          0 allocs/op
BenchmarkPassByValue-10              100                 1.250 ns/op           0 B/op          0 allocs/op
BenchmarkPassByValue-10              100                 1.250 ns/op           0 B/op          0 allocs/op
BenchmarkPassByValue-10              100                 1.250 ns/op           0 B/op          0 allocs/op
BenchmarkPassByReference-10          100                 1.250 ns/op           0 B/op          0 allocs/op
BenchmarkPassByReference-10          100                 1.250 ns/op           0 B/op          0 allocs/op
BenchmarkPassByReference-10          100                 1.250 ns/op           0 B/op          0 allocs/op
BenchmarkPassByReference-10          100                 0.8400 ns/op          0 B/op          0 allocs/op
BenchmarkPassByReference-10          100                 1.660 ns/op           0 B/op          0 allocs/op
```
