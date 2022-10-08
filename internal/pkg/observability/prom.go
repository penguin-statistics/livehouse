package observability

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	ServiceName = "livehouse"
)

var LiveHouseCabinets = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name:    prometheus.BuildFQName(ServiceName, "cabinets", "verify_duration_seconds"),
	Help:    "Duration of report verification in seconds",
	Buckets: prometheus.ExponentialBuckets(0.01, 2, 10),
}, []string{"verifier"})
