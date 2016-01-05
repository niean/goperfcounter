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
BenchmarkMeter   2000000           918 ns/op
BenchmarkMeterMulti  2000000           916 ns/op
BenchmarkGauge   5000000           342 ns/op
BenchmarkGaugeMulti  5000000           333 ns/op
BenchmarkHistogram   2000000           800 ns/op
BenchmarkHistogramMulti  2000000           780 ns/op
ok      github.com/niean/goperfcounter  14.568s

```