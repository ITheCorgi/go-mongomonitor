package mongomonitor

import (
	"sync"
	"sync/atomic"

	"go.mongodb.org/mongo-driver/event"
)

type poolMonitor struct {
	stats stats
	cfg   *config
	mu    *sync.Mutex
}

func NewPoolMonitor(opts ...Option) *event.PoolMonitor {
	cfg := newConfig(opts...)
	pool := &poolMonitor{
		cfg:   cfg,
		stats: stats{},
		mu:    &sync.Mutex{},
	}
	return pool.initConnectionPoolMonitor()
}

func (p *poolMonitor) initConnectionPoolMonitor() *event.PoolMonitor {
	return &event.PoolMonitor{
		Event: func(evt *event.PoolEvent) {
			switch evt.Type {
			case event.GetSucceeded:
				atomic.AddUint64(&p.stats.clientConnections, 1)
				clientConnectionsMetric.Inc()
			case event.ConnectionReturned:
				atomic.AddUint64(&p.stats.clientConnections, ^uint64(0))
				clientConnectionsMetric.Dec()
			}

			if p.cfg.isPoolAlertingOn {
				totalConnections := atomic.LoadUint64(&p.stats.clientConnections)
				clientConnectionUsageMetric.Set(float64(totalConnections) / float64(p.cfg.poolSize) * 100)
			}
		},
	}
}
