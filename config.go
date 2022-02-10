package otelchi

import (
	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/otel/propagation"
	oteltrace "go.opentelemetry.io/otel/trace"
)

// config is used to configure the mux middleware.
type config struct {
	TracerProvider          oteltrace.TracerProvider
	Propagators             propagation.TextMapPropagator
	ChiRoutes               chi.Routes
	RequestMethodInSpanName bool
}

// Option specifies instrumentation configuration options.
type Option interface {
	apply(*config)
}

type optionFunc func(*config)

func (o optionFunc) apply(c *config) {
	o(c)
}

// WithPropagators specifies propagators to use for extracting
// information from the HTTP requests. If none are specified, global
// ones will be used.
func WithPropagators(propagators propagation.TextMapPropagator) Option {
	return optionFunc(func(cfg *config) {
		cfg.Propagators = propagators
	})
}

// WithTracerProvider specifies a tracer provider to use for creating a tracer.
// If none is specified, the global provider is used.
func WithTracerProvider(provider oteltrace.TracerProvider) Option {
	return optionFunc(func(cfg *config) {
		cfg.TracerProvider = provider
	})
}

// WithChiRoutes specified the routes that being used by application. Its main
// purpose is to provide route pattern on span creation.
func WithChiRoutes(routes chi.Routes) Option {
	return optionFunc(func(cfg *config) {
		cfg.ChiRoutes = routes
	})
}

// WithRequestMethodInSpanName is used for adding http request method to span name.
// While this is not necessary for vendors that properly implemented the tracing
// specs (e.g Jaeger, AWS X-Ray, etc...), but for other vendors such as Elastic
// and New Relic this might be helpful.
//
// See following threads for details:
//
// - https://github.com/riandyrn/otelchi/pull/3#issuecomment-1005883910
// - https://github.com/riandyrn/otelchi/issues/6#issuecomment-1034461912
func WithRequestMethodInSpanName(isActive bool) Option {
	return optionFunc(func(cfg *config) {
		cfg.RequestMethodInSpanName = isActive
	})
}
