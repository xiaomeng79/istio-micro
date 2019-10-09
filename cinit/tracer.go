package cinit

import (
	"io"

	"github.com/xiaomeng79/istio-micro/internal/trace"

	"github.com/opentracing/opentracing-go"
	"github.com/xiaomeng79/go-log"
)

var c io.Closer

// 配置文件
func traceInit() {
	// 配置
	c = traceingInit(Config.Trace.Address, Config.Trace.ZipkinURL, Config.Trace.LogTraceSpans, Config.Trace.SamplingRate, Config.Service.Name)
	log.Infof("初始化traceing:%+v", opentracing.GlobalTracer())
}

// 关闭
func tracerClose() {
	if c != nil {
		c.Close()
	}
}

func traceingInit(jaegerURL, zipkinURL string, logTraceSpans bool, samplingRate float64, servicename string) io.Closer {
	cl, err := trace.Configure(servicename, &trace.Options{
		JaegerURL:     jaegerURL,
		ZipkinURL:     zipkinURL,
		LogTraceSpans: logTraceSpans,
		SamplingRate:  samplingRate,
	})
	if err != nil {
		log.Error(err.Error())
	}
	return cl
}
