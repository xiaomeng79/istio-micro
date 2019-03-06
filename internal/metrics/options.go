package metrics

import (
	"github.com/rcrowley/go-metrics"
)

type Option func(c *Options)

type Options struct {
	Registry metrics.Registry
	Prefix   string
}

func Registry(r metrics.Registry) Option {
	return func(o *Options) {
		o.Registry = r
	}
}

func Prefix(p string) Option {
	return func(o *Options) {
		o.Prefix = p
	}
}

func applyOptions(options ...Option) Options {
	opts := Options{
		Registry: metrics.DefaultRegistry,
	}

	for _, option := range options {
		option(&opts)
	}

	return opts
}
