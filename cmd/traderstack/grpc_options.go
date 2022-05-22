package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func appendMetadataGrpcOption(in []grpc.DialOption, key, value string) []grpc.DialOption {
	in = append(in, grpc.WithChainUnaryInterceptor(func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		newCtx := metadata.AppendToOutgoingContext(
			ctx,
			key,
			value,
		)

		return invoker(newCtx, method, req, reply, cc, opts...)
	}))

	in = append(in, grpc.WithChainStreamInterceptor(func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		newCtx := metadata.AppendToOutgoingContext(
			ctx,
			key,
			value,
		)

		return streamer(newCtx, desc, cc, method, opts...)
	}))
	return in
}
