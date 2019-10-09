package trace

import (
	"fmt"
	"io"
	"time"

	"github.com/xiaomeng79/go-log"

	ot "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/transport"
	"github.com/uber/jaeger-client-go/transport/zipkin"
	zk "github.com/uber/jaeger-client-go/zipkin"
)

type holder struct {
	closer io.Closer
	tracer ot.Tracer
}

var (
	httpTimeout = 5 * time.Second
	poolSpans   = jaeger.TracerOptions.PoolSpans(false)
	logger      = spanLogger{}
)

//  indirection for testing
type newZipkin func(url string, options ...zipkin.HTTPOption) (*zipkin.HTTPTransport, error)

//  Configure initializes Istio's tracing subsystem.// //  You typically call this once at process startup.//  Once this call returns, the tracing system is ready to accept data.
func Configure(serviceName string, options *Options) (io.Closer, error) {
	return configure(serviceName, options, zipkin.NewHTTPTransport)
}

func configure(serviceName string, options *Options, nz newZipkin) (io.Closer, error) {
	if err := options.Validate(); err != nil {
		return nil, err
	}

	reporters := make([]jaeger.Reporter, 0, 3)

	sampler, err := jaeger.NewProbabilisticSampler(options.SamplingRate)
	if err != nil {
		return nil, fmt.Errorf("could not build trace sampler: %v", err)
	}

	if options.ZipkinURL != "" {
		trans, err := nz(options.ZipkinURL, zipkin.HTTPLogger(logger), zipkin.HTTPTimeout(httpTimeout))
		if err != nil {
			return nil, fmt.Errorf("could not build zipkin reporter: %v", err)
		}
		reporters = append(reporters, jaeger.NewRemoteReporter(trans))
	}

	if options.JaegerURL != "" {
		reporters = append(reporters, jaeger.NewRemoteReporter(transport.NewHTTPTransport(options.JaegerURL, transport.HTTPTimeout(httpTimeout))))
	}

	if options.LogTraceSpans {
		reporters = append(reporters, logger)
	}

	var rep jaeger.Reporter
	switch len(reporters) {
	case 0:
		return holder{}, nil
	case 1:
		rep = reporters[0]
	default:
		rep = jaeger.NewCompositeReporter(reporters...)
	}

	var tracer ot.Tracer
	var closer io.Closer

	if options.ZipkinURL != "" {
		zipkinPropagator := zk.NewZipkinB3HTTPHeaderPropagator()
		injector := jaeger.TracerOptions.Injector(ot.HTTPHeaders, zipkinPropagator)
		extractor := jaeger.TracerOptions.Extractor(ot.HTTPHeaders, zipkinPropagator)
		tracer, closer = jaeger.NewTracer(serviceName, sampler, rep, poolSpans, injector, extractor, jaeger.TracerOptions.Gen128Bit(true))
	} else {
		tracer, closer = jaeger.NewTracer(serviceName, sampler, rep, poolSpans, jaeger.TracerOptions.Gen128Bit(true))
	}

	//  NOTE: global side effect!
	ot.SetGlobalTracer(tracer)

	return holder{
		closer: closer,
		tracer: tracer,
	}, nil
}

func (h holder) Close() error {
	if ot.GlobalTracer() == h.tracer {
		ot.SetGlobalTracer(ot.NoopTracer{})
	}

	var err error
	if h.closer != nil {
		err = h.closer.Close()
	}

	return err
}

type spanLogger struct{}

//  Report implements the Report() method of jaeger.Reporter
func (spanLogger) Report(span *jaeger.Span) {
	log.Infof("Reporting span operation:%s,span:%s", span.OperationName(), span.String())
}

//  Close implements the Close() method of jaeger.Reporter.
func (spanLogger) Close() {}

//  Error implements the Error() method of log.Logger.
func (spanLogger) Error(msg string) {
	log.Error(msg)
}

//  Infof implements the Infof() method of log.Logger.
func (spanLogger) Infof(msg string, args ...interface{}) {
	log.Infof(msg, args...)
}
