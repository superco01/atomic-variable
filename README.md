# atomic-variable

Run 100 goroutines using atomic increment and standard increment

Sample of race condition result:

`counter atomic result:  100`

`counter result:  99`

---

Benchmarking result:
```
goos: windows
goarch: amd64
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkAtomic-20              1000000000               0.0005023 ns/op
BenchmarkNonAtomic-20           1000000000               0.0005159 ns/op
```