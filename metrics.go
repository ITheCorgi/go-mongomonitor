package mongomonitor

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var namespace string

var (
	clientConnectionsMetric, clientConnectionUsageMetric prometheus.Gauge
)

func initializeConnPoolMetrics() {
	clientConnectionsMetric = promauto.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "mongo_connections_in_pool_count",
			Help:      "number of connections in the MongoDB connection pool",
		})

	clientConnectionUsageMetric = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "mongo_connection_usages_in_percent",
			Help:      "Percentage of connections in use in the MongoDB client connection pool",
		},
	)
}