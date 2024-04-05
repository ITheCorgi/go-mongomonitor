package mongomonitor

import "go.opentelemetry.io/otel/trace"

type Option interface {
	apply(*config)
}

type optionFunc func(*config)

func (o optionFunc) apply(c *config) {
	o(c)
}

// WithTracerProvider specifies a tracer provider to use for creating a tracer.
// If none is specified it is not used
func WithTracerProvider(provider trace.TracerProvider) Option {
	return optionFunc(func(cfg *config) {
		if provider != nil {
			cfg.tracerProvider = provider
			cfg.withTrace = true
		}
	})
}

// WithCommandAttributeDisabled specifies if the MongoDB command is added as an attribute to Spans or not.
// This is disabled by default and the MongoDB command will not be added as an attribute
// to Spans if this option is not provided.
func WithCommandAttributeDisabled(disabled bool) Option {
	return optionFunc(func(cfg *config) {
		cfg.commandAttributeDisabled = disabled
	})
}

func WithScopeName(name string) Option {
	return optionFunc(func(cfg *config) {
		cfg.scopeName = name
	})
}

func WithMetricsEnabled(name string) Option {
	return optionFunc(func(cfg *config) {
		cfg.isMetricsEnabled = true
		cfg.namespace = name
		namespace = name
	})
}

// WithPoolSize specifies a max client connections to mongo
// It is used for alerting in case current connections is near reached this limit
// in case size is not specified, default value of 100 connections is used
// if size=0, it means pool size is not limited and alerting if off
func WithPoolSize(size int) Option {
	return optionFunc(func(cfg *config) {
		cfg.poolSize = size

		if size == 0 {
			cfg.isPoolAlertingOn = false
		}
	})
}
