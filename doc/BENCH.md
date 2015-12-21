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
BenchmarkMeter   2000000           847 ns/op
BenchmarkMeterMulti  2000000           852 ns/op
BenchmarkGauge  10000000           208 ns/op
BenchmarkGaugeMulti 10000000           197 ns/op
BenchmarkGaugeFloat64    5000000           342 ns/op
BenchmarkGaugeFloat64Multi   5000000           354 ns/op
BenchmarkCounter    10000000           193 ns/op
BenchmarkCounterMulti   10000000           201 ns/op
BenchmarkHistogram   2000000           806 ns/op
BenchmarkHistogramMulti  2000000           843 ns/op
BenchmarkTimer   1000000          1585 ns/op
BenchmarkTimerMulti  1000000          1677 ns/op
ok      github.com/niean/goperfcounter  26.706s

```