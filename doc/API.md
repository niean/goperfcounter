API
====

gopfercounter提供了6种类型的统计器，分比为Gauge、GaugeFloat64、Counter、Meter、Histogram、Timer。统计器的含义，参见[java-metrics](http://metrics.dropwizard.io/3.1.0/getting-started/)。


Gauge
----

A gauge metric is an instantaneous reading of a particular value

##### 设置 
+ 接口: Gauge(name string, value int64)
+ 参数: value - 记录的数值
+ 例子:

```go
Gauge("queueSize", int64(13))
```

##### 设置 
+ 接口: SetGaugeValue(name string, value float64)
+ Alias: GaugeFloat64(name string, value float64)
+ 参数: value - 记录的数值
+ 例子:

```go
GaugeFloat64("requestRate", float64(13.14))
SetGaugeValue("requestRate", float64(13.14))
```

##### 获取
+ 接口: GetGaugeValue(name string) float64
+ 例子:

```go
reqRate := GetGaugeValue("requestRate")
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

##### 获取变化率
+ 接口: GetMeterRate(name string) float64
+ 例子:

```go
// pv发生次数的时间平均，单位CPS。计时范围为，本接口两次调用的时间差。
pvRateMean := GetMeterRate("pageView")
```

##### 获取累计的平均值
+ 接口: GetMeterRateMean(name string) float64
+ 例子:

```go
// pv发生次数的时间平均，单位CPS。计时范围为，goperfcounter完成初始，至当前时刻。
pvRateMean := GetMeterRateMean("pageView")
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

##### 获取最大值
+ 接口: GetHistogramMax(name string) int64
+ 例子:

```go
max := GetHistogramMax("processNum")
```

##### 获取最小值
+ 接口: GetHistogramMin(name string) int64
+ 例子:

```go
min := GetHistogramMin("processNum")
```

##### 获取平均值
+ 接口: GetHistogramMean(name string) float64
+ 例子:

```go
mean := GetHistogramMean("processNum")
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

