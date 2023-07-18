package metrics

import "github.com/prometheus/client_golang/prometheus"

type MetricInit interface {
	Init() []prometheus.Collector
}

var metrics = []MetricInit{
	Http,
}
