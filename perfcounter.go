package goperfcounter

import (
	"strings"
	"time"

	"github.com/niean/go-metrics-lite"
)

func init() {
	// init cfg
	err := loadConfig()
	if err != nil {
		setDefaultConfig()
	}
	cfg := config()

	// init http
	if cfg.Http.Enabled {
		go startHttp(cfg.Http.Listen, cfg.Debug)
	}

	// base collector cron
	if len(cfg.Bases) > 0 {
		go collectBase(cfg.Bases)
	}

	// push cron
	if cfg.Push.Enabled {
		go pushToFalcon()
	}
}

// counter
func Counter(name string, count int64) {
	SetCounterCount(name, count)
}
func SetCounterCount(name string, count int64) {
	rr := gpCounter.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Counter); ok {
			r.Inc(count)
		}
		return
	}

	r := metrics.NewCounter()
	r.Inc(count)
	if err := gpCounter.Register(name, r); isDuplicateMetricError(err) {
		r := gpCounter.Get(name).(metrics.Counter)
		r.Inc(count)
	}
}

func GetCounterCount(name string) int64 {
	rr := gpCounter.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Counter); ok {
			return r.Count()
		}
	}
	return 0
}

// gauge
func Gauge(name string, value int64) {
	SetGaugeValue(name, value)
}
func SetGaugeValue(name string, value int64) {
	rr := gpGauge.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Gauge); ok {
			r.Update(value)
		}
		return
	}

	r := metrics.NewGauge()
	r.Update(value)
	if err := gpGauge.Register(name, r); isDuplicateMetricError(err) {
		r := gpGauge.Get(name).(metrics.Gauge)
		r.Update(value)
	}
}

func GetGaugeValue(name string) int64 {
	rr := gpGauge.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Gauge); ok {
			return r.Value()
		}
	}
	return 0
}

// gaugefloat64
func GaugeFloat64(name string, value float64) {
	SetGaugeFloat64Value(name, value)
}
func SetGaugeFloat64Value(name string, value float64) {
	rr := gpGaugeFloat64.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.GaugeFloat64); ok {
			r.Update(value)
		}
		return
	}

	r := metrics.NewGaugeFloat64()
	r.Update(value)
	if err := gpGaugeFloat64.Register(name, r); isDuplicateMetricError(err) {
		r := gpGaugeFloat64.Get(name).(metrics.GaugeFloat64)
		r.Update(value)
	}
}

func GetGaugeFloat64Value(name string) float64 {
	rr := gpGaugeFloat64.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.GaugeFloat64); ok {
			return r.Value()
		}
	}
	return 0.0
}

// meter
func Meter(name string, count int64) {
	SetMeterCount(name, count)
}
func SetMeterCount(name string, count int64) {
	rr := gpMeter.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Meter); ok {
			r.Mark(count)
		}
		return
	}

	r := metrics.NewMeter()
	r.Mark(count)
	if err := gpMeter.Register(name, r); isDuplicateMetricError(err) {
		r := gpMeter.Get(name).(metrics.Meter)
		r.Mark(count)
	}
}

func GetMeterCount(name string) int64 {
	rr := gpMeter.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Meter); ok {
			return r.Count()
		}
	}
	return 0
}

func GetMeterRateMean(name string) float64 {
	rr := gpMeter.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Meter); ok {
			return r.RateMean()
		}
	}
	return 0.0
}

func GetMeterRateStep(name string) float64 {
	rr := gpMeter.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Meter); ok {
			return r.RateStep()
		}
	}
	return 0.0
}

func GetMeterRate1(name string) float64 {
	rr := gpMeter.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Meter); ok {
			return r.Rate1()
		}
	}
	return 0.0
}

func GetMeterRate5(name string) float64 {
	rr := gpMeter.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Meter); ok {
			return r.Rate5()
		}
	}
	return 0.0
}

func GetMeterRate15(name string) float64 {
	rr := gpMeter.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Meter); ok {
			return r.Rate15()
		}
	}
	return 0.0
}

// histogram
func Histogram(name string, count int64) {
	SetHistogramCount(name, count)
}
func SetHistogramCount(name string, count int64) {
	rr := gpHistogram.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Histogram); ok {
			r.Update(count)
		}
		return
	}

	s := metrics.NewExpDecaySample(1028, 0.015)
	r := metrics.NewHistogram(s)
	r.Update(count)
	if err := gpHistogram.Register(name, r); isDuplicateMetricError(err) {
		r := gpHistogram.Get(name).(metrics.Histogram)
		r.Update(count)
	}
}

func GetHistogramCount(name string) int64 {
	rr := gpHistogram.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Histogram); ok {
			return r.Count()
		}
	}
	return 0
}
func GetHistogramMax(name string) int64 {
	rr := gpHistogram.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Histogram); ok {
			return r.Max()
		}
	}
	return 0
}
func GetHistogramMin(name string) int64 {
	rr := gpHistogram.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Histogram); ok {
			return r.Min()
		}
	}
	return 0
}
func GetHistogramSum(name string) int64 {
	rr := gpHistogram.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Histogram); ok {
			return r.Sum()
		}
	}
	return 0
}
func GetHistogramMean(name string) float64 {
	rr := gpHistogram.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Histogram); ok {
			return r.Mean()
		}
	}
	return 0.0
}
func GetHistogramStdDev(name string) float64 {
	rr := gpHistogram.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Histogram); ok {
			return r.StdDev()
		}
	}
	return 0.0
}
func GetHistogram50th(name string) float64 {
	rr := gpHistogram.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Histogram); ok {
			return r.Percentile(0.5)
		}
	}
	return 0.0
}
func GetHistogram75th(name string) float64 {
	rr := gpHistogram.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Histogram); ok {
			return r.Percentile(0.75)
		}
	}
	return 0.0
}
func GetHistogram95th(name string) float64 {
	rr := gpHistogram.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Histogram); ok {
			return r.Percentile(0.95)
		}
	}
	return 0.0
}
func GetHistogram99th(name string) float64 {
	rr := gpHistogram.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Histogram); ok {
			return r.Percentile(0.99)
		}
	}
	return 0.0
}
func GetHistogram999th(name string) float64 {
	rr := gpHistogram.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Histogram); ok {
			return r.Percentile(0.999)
		}
	}
	return 0.0
}

// timer
func Timer(name string, ms int64) {
	SetTimerCount(name, ms)
}
func SetTimerCount(name string, ms int64) {
	count := time.Duration(ms) * time.Millisecond
	rr := gpTimer.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Timer); ok {
			r.Update(count)
		}
		return
	}

	r := metrics.NewTimer()
	r.Update(count)
	if err := gpTimer.Register(name, r); isDuplicateMetricError(err) {
		r := gpTimer.Get(name).(metrics.Timer)
		r.Update(count)
	}
}

func GetTimerCount(name string) int64 {
	rr := gpTimer.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Timer); ok {
			return r.Count()
		}
	}
	return 0
}
func GetTimerMax(name string) int64 {
	rr := gpTimer.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Timer); ok {
			return r.Max()
		}
	}
	return 0
}
func GetTimerMin(name string) int64 {
	rr := gpTimer.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Timer); ok {
			return r.Min()
		}
	}
	return 0
}
func GetTimerMean(name string) float64 {
	rr := gpTimer.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Timer); ok {
			return r.Mean()
		}
	}
	return 0.0
}

func GetTimerStdDev(name string) float64 {
	rr := gpTimer.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Timer); ok {
			return r.StdDev()
		}
	}
	return 0.0
}
func GetTimer50th(name string) float64 {
	rr := gpTimer.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Timer); ok {
			return r.Percentile(0.5)
		}
	}
	return 0.0
}
func GetTimer75th(name string) float64 {
	rr := gpTimer.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Timer); ok {
			return r.Percentile(0.75)
		}
	}
	return 0.0
}
func GetTimer95th(name string) float64 {
	rr := gpTimer.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Timer); ok {
			return r.Percentile(0.95)
		}
	}
	return 0.0
}
func GetTimer99th(name string) float64 {
	rr := gpTimer.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Timer); ok {
			return r.Percentile(0.99)
		}
	}
	return 0.0
}
func GetTimer999th(name string) float64 {
	rr := gpTimer.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Timer); ok {
			return r.Percentile(0.999)
		}
	}
	return 0.0
}
func GetTimerRateMean(name string) float64 {
	rr := gpTimer.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Timer); ok {
			return r.RateMean()
		}
	}
	return 0.0
}
func GetTimerRate1(name string) float64 {
	rr := gpTimer.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Timer); ok {
			return r.Rate1()
		}
	}
	return 0.0
}
func GetTimerRate5(name string) float64 {
	rr := gpTimer.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Timer); ok {
			return r.Rate5()
		}
	}
	return 0.0
}
func GetTimerRate15(name string) float64 {
	rr := gpTimer.Get(name)
	if rr != nil {
		if r, ok := rr.(metrics.Timer); ok {
			return r.Rate15()
		}
	}
	return 0.0
}

// internal
func isDuplicateMetricError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Index(err.Error(), "duplicate metric:") == 0
}
