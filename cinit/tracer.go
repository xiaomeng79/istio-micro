package cinit

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"github.com/xiaomeng79/go-log"
	"io"
)

var c io.Closer

//配置文件
func traceInit() {
	//配置
	c = traceingInit(Config.Trace.Address, Config.Service.Name)
	log.Infof("初始化traceing:%+v", opentracing.GlobalTracer())
}

//关闭
func tracerClose() {
	if c != nil {
		c.Close()
	}
}

func traceingInit(address, servicename string) io.Closer {
	//配置
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	metricsFactory := metrics.NewLocalFactory(0)
	_metrics := jaeger.NewMetrics(metricsFactory, nil)

	sender, err := jaeger.NewUDPTransport(address, 0)
	if err != nil {
		log.Info("could not initialize jaeger sender: " + err.Error())
		return nil
	}

	repoter := jaeger.NewRemoteReporter(sender, jaeger.ReporterOptions.Metrics(_metrics))

	// Initialize tracer with a logger and a metrics factory
	closer, err := cfg.InitGlobalTracer(
		servicename,
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
		jaegercfg.Reporter(repoter),
	)

	if err != nil {
		log.Info("could not initialize jaeger tracer: " + err.Error())
		return nil
	}
	return closer
}
