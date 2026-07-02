package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	Requests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "gateway_requests_total",
			Help: "Total number of requests",
		},
	)

	Throttled = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "gateway_throttled_total",
			Help: "Total throttled requests",
		},
	)
)

func Register() {

	prometheus.MustRegister(Requests)

	prometheus.MustRegister(Throttled)

}
