package goperfcounter

import (
	"time"

	"github.com/niean/go-metrics-lite"
	"github.com/niean/goperfcounter/base"
)

var (
	gpCounter      = metrics.NewRegistry()
	gpGauge        = metrics.NewRegistry()
	gpGaugeFloat64 = metrics.NewRegistry()
	gpMeter        = metrics.NewRegistry()
	gpHistogram    = metrics.NewRegistry()
	gpTimer        = metrics.NewRegistry()
	gpDebug        = metrics.NewRegistry()
	gpRuntime      = metrics.NewRegistry()
	values         = make(map[string]metrics.Registry) //readonly,mappings of metrics
)

func init() {
	values["counter"] = gpCounter
	values["gauge"] = gpGauge
	values["gauge.float64"] = gpGaugeFloat64
	values["meter"] = gpMeter
	values["histogram"] = gpHistogram
	values["timer"] = gpTimer
	values["debug"] = gpDebug
	values["runtime"] = gpRuntime
}

//
func rawMetric(mtype string) map[string]interface{} {
	data := make(map[string]interface{})
	if v, ok := values[mtype]; ok {
		data[mtype] = v.Values()
	}
	return data
}

func rawMetrics() map[string]interface{} {
	data := make(map[string]interface{})
	for key, v := range values {
		data[key] = v.Values()
	}
	return data
}

func collectBase(bases []string) {
	// start base collect after 30sec
	time.Sleep(time.Duration(30) * time.Second)

	if contains(bases, "debug") {
		base.RegisterAndCaptureDebugGCStats(gpDebug, 5e9)
	}

	if contains(bases, "runtime") {
		base.RegisterAndCaptureRuntimeMemStats(gpRuntime, 5e9)
	}
}

func contains(bases []string, name string) bool {
	for _, n := range bases {
		if n == name {
			return true
		}
	}
	return false
}
