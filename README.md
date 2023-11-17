# Go Performance
Benchmark the various cases by using value, pointer.

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