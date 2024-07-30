package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "request",
			Subsystem: "gRPC",
			Name:      "http_requests_total",
			Help:      "Total number of gRPC requests.",
		},
		[]string{"grpc_call"})

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "request",
			Subsystem: "gRPC",
			Name:      "http_request_duration_seconds",
			Help:      "Duration of gRPC requests.",
		},
		[]string{"grpc_call"})
)
