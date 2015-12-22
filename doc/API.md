API
====

gopfercounter提供了6种类型的统计器，分比为Gauge、GaugeFloat64、Counter、Meter、Histogram、Timer。统计器的含义，参见[java-metrics](http://metrics.dropwizard.io/3.1.0/getting-started/)。


Gauge
----

A gauge metric is an instantaneous reading of a particular int64 value

##### 设置 
+ 接口: SetGaugeValue(name string, value int64)
+ Alias: Gauge(name string, value int64)
+ 参数: value - 记录的数值
+ 例子:

```go
SetGaugeValue("queueSize", int64(1))
```

##### 获取
+ 接口: GetGaugeValue(name string) int64
+ 例子:

```go
qsize := GetCounterCount("queueSize")
```

GaugeFloat64
----
A gauge metric is an instantaneous reading of a particular float64 value

##### 设置 
+ 接口: SetGaugeFloat64Value(name string, value float64)
+ Alias: GaugeFloat64(name string, value float64)
+ 参数: value - 记录的数值
+ 例子:

```go
SetGaugeFloat64Value("requestRate", float64(19.86))
```

##### 获取
+ 接口: GetGaugeFloat64Value(name string) float64
+ 例子:

```go
rRate := GetGaugeFloat64Value("requestRate")
```

Counter
----

An incrementing and decrementing gauge metric

##### 设置 
+ 接口: SetCounterCount(name string, count int64)
+ Alias: Counter(name string, count int64)
+ 参数: count - 记录该值的相对变化量
+ 例子:

```go
// 新增一条链接
SetCounterCount("ConnectionNum", int64(1))
// 断开一条链接
SetCounterCount("ConnectionNum", int64(-1))
```

##### 获取累计的值
+ 接口: GetCounterCount(name string) int64
+ 例子:

```go
// 获取当前链接数
cNum := GetCounterCount("ConnectionNum")
```

Meter
----

A meter metric which measures mean throughput and one-, five-, and fifteen-minute exponentially-weighted moving average throughputs.

关于EWMA, 点击[这里](http://en.wikipedia.org/wiki/Moving_average#Exponential_moving_average)

##### 设置 
+ 接口: SetMeterCount(name string, value int64)
+ Alias: Meter(name string, value int64)
+ 参数: value - 该事件发生的次数 	
+ 例子:

```go
// 页面访问次数统计，每来一次访问，pv加1
SetMeterCount("pageView", int64(1))
```

##### 获取累计的值
+ 接口: GetMeterCount(name string) int64
+ 例子:

```go
// 获取pv次数的总和
pvSum := GetMeterCount("pageView")
```

##### 获取累计的平均值
+ 接口: GetMeterRateMean(name string) float64
+ 例子:

```go
// pv发生次数的时间平均，单位CPS。计时起点，为goperfcounter完成初始化。
pvRateMean := GetMeterRateMean("pageView")
```

##### 获取累计的平均值
+ 接口: GetMeterRateStep(name string) float64
+ 例子:

```go
// 一个MeterTicker内，pv发生次数的时间平均，单位CPS。MeterTicker周期为5s。
pvRateStep := GetMeterRateStep("pageView")
```

##### 获取1min的滑动平均
+ 接口: GetMeterRate1(name string) float64
+ 例子:

```go
// pv发生次数的1min滑动平均值，单位CPS
pvRate1Min := GetMeterRate1("pageView")
```

##### 获取5min的滑动平均
+ 接口: GetMeterRate5(name string) float64
+ 例子:

```go
// pv发生次数的5min滑动平均值，单位CPS
pvRate5Min := GetMeterRate5("pageView")
```

##### 获取15min的滑动平均
+ 接口: GetMeterRate15(name string) float64
+ 例子:

```go
// pv发生次数的15min滑动平均值，单位CPS
pvRate15Min := GetMeterRate15("pageView")
```

Histogram
----

A histogram measures the [statistical distribution](http://www.johndcook.com/standard_deviation.html) of values in a stream of data. In addition to minimum, maximum, mean, etc., it also measures median, 75th, 90th, 95th, 98th, 99th, and 99.9th percentiles

##### 设置 
+ 接口: SetHistogramCount(name string, count int64)
+ Alias: Histogram(name string, count int64)
+ 参数: count - 该记录当前采样点的取值
+ 例子:

```go
// 设置当前同时处理请求的并发度
SetHistogramCount("processNum", int64(325))
```

##### 获取中值 50thPecentile
+ 接口: GetHistogram50th(name string) float64
+ 例子:

```go
// 获取并发度的中位数
pNum50th := GetHistogram50th("processNum")
```

##### 获取75thPecentile
+ 接口: GetHistogram75th(name string) float64
+ 例子:

```go
// 获取所有采样数据中，处于75%的并发度
pNum75th := GetHistogram75th("processNum")
```

##### 获取95thPecentile
+ 接口: GetHistogram95th(name string) float64
+ 例子:

```go
// 获取所有采样数据中，处于95%的并发度
pNum95th := GetHistogram95th("processNum")
```

##### 获取99thPecentile
+ 接口: GetHistogram99th(name string) float64
+ 例子:

```go
// 获取所有采样数据中，处于99%的并发度
pNum99th := GetHistogram99th("processNum")
```

##### 获取999thPecentile
+ 接口: GetHistogram999th(name string) float64
+ 例子:

```go
// 获取所有采样数据中，处于99.9%的并发度
pNum999th := GetHistogram999th("processNum")
```

Timer
----

A timer metric which aggregates timing durations and provides duration statistics, plus throughput statistics via [Meter](###Meter)

##### 设置 
+ 接口: SetTimerCount(name string, ms int64)
+ Alias: Timer(name string, ms int64)
+ 参数: ms - 记录该事件的持续时间，单位ms
+ 例子:

```go
// 调用getUserId方法执行的时间, 20ms.
SetTimerCount("timeOfGetUserId", 20)
```

##### 获取平均值
+ 接口: GetTimerMean(name string) float64
+ 例子:

```go
mean := GetTimerMean("timeOfGetUserId")
```

##### 获取最大值
+ 接口: GetTimerMax(name string) int64
+ 例子:

```go
max := GetTimerMax("timeOfGetUserId")
```

##### 获取最小值
+ 接口: GetTimerMin(name string) int64
+ 例子:

```go
min := GetTimerMin("timeOfGetUserId")
```

