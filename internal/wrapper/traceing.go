package wrapper

import (
	"context"
	"github.com/xiaomeng79/go-log"
	"github.com/xiaomeng79/istio-micro/internal/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TraceingingUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	log.Infof("md:%+v", md, ctx)
	log.Infof("ctx:%+v", ctx, ctx)
	resp, err := handler(ctx, req)
	log.Infof("gRPC method: %s, resp: %v", info.FullMethod, resp, ctx)
	return resp, err
}

func TraceingUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(parentCtx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		parentCtx = trace.TraceToRpcHeader(parentCtx)
		err := invoker(parentCtx, method, req, reply, cc, opts...)
		return err
	}
}
