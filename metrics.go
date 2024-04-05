package mongomonitor

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var namespace string

var (
	clientConnectionsMetric = promauto.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "client_connections_in_mongo_pool_count",
			Help:      "number of connections in the MongoDB connection pool",
		})

	clientConnectionUsageMetric = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "client_connection_mongo_usage_in_percent",
			Help:      "Percentage of connections in use in the MongoDB client connection pool",
		},
	)
)

func init() {
	prometheus.MustRegister(clientConnectionsMetric)
	prometheus.MustRegister(clientConnectionUsageMetric)
}
