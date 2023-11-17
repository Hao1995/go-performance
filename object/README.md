
# Object

## Native
Edit person object: `Person v.s. *Person`

```
go test -bench=. ./object/native/ -benchmem -count=1 -benchtime=100x
```

```
BenchmarkPassByValue-10              100                 1.250 ns/op           0 B/op          0 allocs/op
BenchmarkPassByReference-10          100                 2.080 ns/op           0 B/op          0 allocs/op
```
