package logger

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/infra/tracing"
	"google.golang.org/grpc"
	"time"
)

func LoggingInterceptor(log domain.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = tracing.AddRequestIdGrpc(ctx)
		t := time.Now()
		log.Info(ctx).
			Str("method", info.FullMethod).
			Interface("req", req).
			Time("time", t).
			Msg("starting gRPC call")

		result, err := handler(ctx, req)

		log.Info(ctx).
			Str("method", info.FullMethod).
			Interface("result", result).
			Msg("finished gRPC call")
		return result, err
	}
}
