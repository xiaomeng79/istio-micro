package trace

import (
	"context"
	"google.golang.org/grpc/metadata"
	"net/http"
)

const (
	IstioRequestId    = "x-request-id"
	IstioTraceId      = "x-b3-traceid"
	IstioSpanId       = "x-b3-spanid"
	IstioParentSpanId = "x-b3-parentspanid"
	IstioSampled      = "x-b3-sampled"
	IstioFlags        = "x-b3-flags"
	IstioSpanContext  = "x-ot-span-context"
)

var otHeaders = []string{
	IstioRequestId,
	IstioTraceId,
	IstioSpanId,
	IstioParentSpanId,
	IstioSampled,
	IstioFlags,
	IstioSpanContext,
}

const (
	Log_Trace = "log_trace"
)

func TraceFromHttpHeader(ctx context.Context, header http.Header) context.Context {
	md := metadata.New(make(map[string]string))
	for _, v := range otHeaders {
		if rh := header.Get(v); len(rh) > 0 {
			md.Set(v, rh)
		}
	}
	if len(header.Get(IstioTraceId)) > 0 {
		ctx = context.WithValue(ctx, Log_Trace, header.Get(IstioTraceId)+":"+header.Get(IstioSpanId)+":"+header.Get(IstioParentSpanId))
	}
	return metadata.NewIncomingContext(ctx, md)
}

func TraceFromRpcHeader(ctx context.Context) context.Context {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		tid := md.Get(IstioTraceId)
		sid := md.Get(IstioSpanId)
		pid := md.Get(IstioParentSpanId)
		if len(tid) > 0 {
			ctx = context.WithValue(ctx, Log_Trace, tid[0]+":"+sid[0]+":"+pid[0])
		}
	}
	return ctx
}

func TraceToRpcHeader(ctx context.Context) context.Context {
	md, _ := metadata.FromIncomingContext(ctx)
	return metadata.NewOutgoingContext(ctx, md)
}

//func TraceToHttpHeader(ctx context.Context,)
