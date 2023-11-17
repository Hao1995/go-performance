# Slice

## Native
Edit the slice content: `[]People v.s. []*People`

```
go test -bench=. ./slice/native/ -benchmem -count=1 -benchtime=100x
```

```
BenchmarkPassByValue-10              100                51.66 ns/op            0 B/op          0 allocs/op
BenchmarkPassByReference-10          100                42.50 ns/op            0 B/op          0 allocs/op
```

## Convert Slice

Convert people slice to member slice: `[]People -> []Member v.s. []*People -> []*Member`

```
go test -bench=. ./slice/conv_slice/ -benchmem -count=1 -benchtime=100x
```

```
BenchmarkPassByValue-10              100               142.1 ns/op           896 B/op          1 allocs/op
BenchmarkPassByReference-10          100               774.6 ns/op          1120 B/op         21 allocs/op
```

## New Slice

Get the edited result by return new slice when using slice of value: `[]People -> []People v.s. []*People`

```
go test -bench=. ./slice/new_slice/ -benchmem -count=1 -benchtime=100x
```

```
BenchmarkPassByValue-10              100               525.0 ns/op           896 B/op          1 allocs/op
BenchmarkPassByReference-10          100              1190 ns/op             969 B/op         20 allocs/op

```
## Change Size

If append an item to a slice of value, all data will be copy once: `[]People v.s. []*People`

```
go test -bench=. ./slice/change_size/ -benchmem -count=1 -benchtime=100x
```

```
BenchmarkPassByValue-10              100               338.3 ns/op          1792 B/op          1 allocs/op
BenchmarkPassByReference-10          100               144.2 ns/op           377 B/op          2 allocs/op

```