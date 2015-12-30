Bench Test
====

Run
----

```bash
cd $GOPATH/src/github.com/niean/goperfcounter && go test -test.bench=".*"

```

Result
----

```
PASS
BenchmarkMeter   2000000           892 ns/op
BenchmarkMeterMulti  2000000           873 ns/op
BenchmarkGauge   5000000           338 ns/op
BenchmarkGaugeMulti  5000000           350 ns/op
BenchmarkCounter    10000000           186 ns/op
BenchmarkCounterMulti   10000000           178 ns/op
BenchmarkHistogram   2000000           758 ns/op
BenchmarkHistogramMulti  2000000           751 ns/op
ok      github.com/niean/goperfcounter  18.383s

```