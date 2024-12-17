package middleware

import (
	"context"

	"github.com/c2micro/c2mshr/defaults"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryClientInterceptor(t string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx = metadata.AppendToOutgoingContext(ctx, defaults.GrpcAuthOperatorHeader, t)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func StreamClientInterceptor(t string) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		ctx = metadata.AppendToOutgoingContext(ctx, defaults.GrpcAuthOperatorHeader, t)
		return streamer(ctx, desc, cc, method, opts...)
	}
}
