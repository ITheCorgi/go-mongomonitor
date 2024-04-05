package mongomonitor

import (
	"sync"

	"go.opentelemetry.io/otel/trace"
)

type config struct {
	tracerProvider           trace.TracerProvider
	tracer                   trace.Tracer
	scopeName                string
	namespace                string
	poolSize                 int
	withTrace                bool
	isMetricsEnabled         bool
	isPoolAlertingOn         bool
	commandAttributeDisabled bool
}

func newConfig(opts ...Option) config {
	cfgOnce := sync.Once{}
	cfg := config{}

	cfgOnce.Do(func() {
		cfg = getDefaultConfig()

		for _, opt := range opts {
			opt.apply(&cfg)
		}

		if cfg.withTrace {
			cfg.tracer = cfg.tracerProvider.Tracer(
				cfg.scopeName,
				trace.WithInstrumentationVersion("0.49.0"),
			)
		}
	})

	return cfg
}

func getDefaultConfig() config {
	const defaultConnectionPoolSize = 100

	return config{
		poolSize:                 defaultConnectionPoolSize,
		isPoolAlertingOn:         true,
		commandAttributeDisabled: true,
	}
}

type stats struct {
	clientConnections uint64
}
