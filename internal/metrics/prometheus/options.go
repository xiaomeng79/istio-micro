package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

var defaultMetricPath = "/metrics"

type Option func(c *Options)

type Options struct {
	Registry    prometheus.Registerer
	MetricsPath string
	Namespace   string
	Subsystem   string
}

func Registry(r prometheus.Registerer) Option {
	return func(o *Options) {
		o.Registry = r
	}
}

func MetricsPath(v string) Option {
	return func(o *Options) {
		o.MetricsPath = v
	}
}

func Namespace(v string) Option {
	return func(o *Options) {
		o.Namespace = v
	}
}

func Subsystem(v string) Option {
	return func(o *Options) {
		o.Subsystem = v
	}
}

func applyOptions(options ...Option) Options {
	opts := Options{
		Registry:    prometheus.DefaultRegisterer,
		MetricsPath: defaultMetricPath,
	}

	for _, option := range options {
		option(&opts)
	}

	return opts
}
