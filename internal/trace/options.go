package trace

import (
	"errors"
)

//  Options defines the set of options supported by Istio's component tracing package.
type Options struct {
	//  URL of zipkin collector (example: 'http://zipkin:9411/api/v1/spans').
	ZipkinURL string

	//  URL of jaeger HTTP collector (example: 'http://jaeger:14268/api/traces?format=jaeger.thrift').
	JaegerURL string

	//  Whether or not to emit trace spans as log records.
	LogTraceSpans bool

	//  SamplingRate controls the rate at which a process will decide to generate trace spans.
	SamplingRate float64
}

//  DefaultOptions returns a new set of options, initialized to the defaults
func DefaultOptions() *Options {
	return &Options{}
}

//  Validate returns whether the options have been configured correctly or an error
func (o *Options) Validate() error {
	//  due to a race condition in the OT libraries somewhere, we can't have both tracing outputs active at once
	if o.JaegerURL != "" && o.ZipkinURL != "" {
		return errors.New("can't have Jaeger and Zipkin outputs active simultaneously")
	}

	if o.SamplingRate > 1.0 || o.SamplingRate < 0.0 {
		return errors.New("sampling rate must be in the range: [0.0, 1.0]")
	}

	return nil
}

//  TracingEnabled returns whether the given options enable tracing to take place.
func (o *Options) TracingEnabled() bool {
	return o.JaegerURL != "" || o.ZipkinURL != "" || o.LogTraceSpans
}
