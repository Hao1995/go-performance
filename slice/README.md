# Slice

## Native
Edit the slice content: `[]People v.s. []*People`

```
go test -bench=. ./slice/native/ -benchmem -count=1 -benchtime=100x
```

```
BenchmarkPassByValue-10              100           2792550 ns/op          400015 B/op          0 allocs/op
BenchmarkPassByReference-10          100           2185102 ns/op          560056 B/op      10000 allocs/op
```

## Convert Slice

Convert people slice to member slice: `[]People -> []Member v.s. []*People -> []*Member`

```
go test -bench=. ./slice/conv_slice/ -benchmem -count=1 -benchtime=100x
```

```
BenchmarkPassByValue-10              100          19862205 ns/op        40401617 B/op          1 allocs/op
BenchmarkPassByReference-10          100          28915872 ns/op        56563663 B/op    1010001 allocs/op
```

## New Slice

Get the edited result by return new slice when using slice of value: `[]People -> []People v.s. []*People`

```
go test -bench=. ./slice/new_slice/ -benchmem -count=1 -benchtime=100x
```

```
BenchmarkPassByValue-10              100          20422887 ns/op        40401571 B/op          1 allocs/op
BenchmarkPassByReference-10          100          23689532 ns/op        48560089 B/op    1010000 allocs/op

```
## Change Size

If append an item to a slice of value, all data will be copy once: `[]People v.s. []*People`

```
go test -bench=. ./slice/change_size/ -benchmem -count=1 -benchtime=100x
```

```
BenchmarkPassByValue-10              100          17640877 ns/op        50412176 B/op          1 allocs/op
BenchmarkPassByReference-10          100           3990002 ns/op        10562536 B/op      10002 allocs/op

```