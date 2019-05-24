package trace

//
//
//import (
//	"context"
//	"github.com/opentracing/opentracing-go"
//	"google.golang.org/grpc/metadata"
//	"net/http"
//)
//
////istio
//
//const (
//	prefixTracerState  = "x-b3-"
//	zipkinTraceID      = prefixTracerState + "traceid"
//	zipkinSpanID       = prefixTracerState + "spanid"
//	zipkinParentSpanID = prefixTracerState + "parentspanid"
//	zipkinSampled      = prefixTracerState + "sampled"
//	zipkinFlags        = prefixTracerState + "flags"
//)
//
//var otHeaders = []string{
//	zipkinTraceID,
//	zipkinSpanID,
//	zipkinParentSpanID,
//	zipkinSampled,
//	zipkinFlags,
//	}
//
//
//
////记录tag
//func tag(ctx context.Context, sp opentracing.Span) (context.Context, opentracing.Span) {
//	ctx = opentracing.ContextWithSpan(ctx, sp)
//
//	//加tag
//	//s := strings.Split(fmt.Sprintf("%v", sp), ":")
//	//if len(s) >= 3 {
//	//	sp.SetTag("trace_id", s[0])
//	//	sp.SetTag("span_id", s[1])
//	//	sp.SetTag("parent_id", s[2])
//	//}
//	return ctx, sp
//}
//
//func traceIntoContextByGlobalTracer(ctx context.Context, tracer opentracing.Tracer, name string) (context.Context, opentracing.Span, error) {
//	md, ok := metadata.FromIncomingContext(ctx)
//	if !ok {
//		md = make(map[string]string)
//	}
//	var sp opentracing.Span
//	wireContext, err := tracer.Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
//	if err != nil {
//		sp = tracer.StartSpan(name)
//	} else {
//		sp = tracer.StartSpan(name, opentracing.ChildOf(wireContext))
//	}
//	if err := sp.Tracer().Inject(sp.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md)); err != nil {
//		return nil, nil, err
//	}
//	ctx, sp = tag(ctx, sp)
//	ctx = metadata.NewContext(ctx, md)
//	return ctx, sp, nil
//}
//
//func traceFromHeaderByGlobalTracer(ctx context.Context, tracer opentracing.Tracer, name string, header http.Header) (context.Context, opentracing.Span, error) {
//	var sp opentracing.Span
//	wireContext, err := tracer.Extract(opentracing.TextMap, opentracing.HTTPHeadersCarrier(header))
//	if err != nil {
//		sp = tracer.StartSpan(name)
//	} else {
//		sp = tracer.StartSpan(name, opentracing.ChildOf(wireContext))
//	}
//	md, ok := metadata.FromContext(ctx)
//	if !ok {
//		md = make(map[string]string)
//	}
//	if err := sp.Tracer().Inject(sp.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md)); err != nil {
//		return nil, nil, err
//	}
//	ctx, sp = tag(ctx, sp)
//	ctx = metadata.NewContext(ctx, md)
//	return ctx, sp, nil
//}
//
//func traceToHeaderByGlobalTracer(ctx context.Context, tracer opentracing.Tracer, name string, header http.Header) (context.Context, opentracing.Span, error) {
//	md, ok := metadata.FromContext(ctx)
//	if !ok {
//		md = make(map[string]string)
//	}
//	var sp opentracing.Span
//	wireContext, err := tracer.Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
//	if err != nil {
//		sp = tracer.StartSpan(name)
//	} else {
//		sp = tracer.StartSpan(name, opentracing.ChildOf(wireContext))
//	}
//	if err := sp.Tracer().Inject(sp.Context(), opentracing.TextMap, opentracing.HTTPHeadersCarrier(header)); err != nil {
//		return nil, nil, err
//	}
//	ctx, sp = tag(ctx, sp)
//	return ctx, sp, nil
//}
//
////opentracing从context获取,写入context，适用RPC
//func TraceIntoContext(ctx context.Context, name string) (context.Context, opentracing.Span, error) {
//	return traceIntoContextByGlobalTracer(ctx, opentracing.GlobalTracer(), name)
//}
//
////opentracing从header获取,写入context,适用获取http
//func TraceFromHeader(ctx context.Context, name string, header http.Header) (context.Context, opentracing.Span, error) {
//	return traceFromHeaderByGlobalTracer(ctx, opentracing.GlobalTracer(), name, header)
//}
//
////opentracing从context获取,写入http,适用将调用http
//func TraceToHeader(ctx context.Context, name string, header http.Header) (context.Context, opentracing.Span, error) {
//	return traceToHeaderByGlobalTracer(ctx, opentracing.GlobalTracer(), name, header)
//}
