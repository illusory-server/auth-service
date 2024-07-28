package logger

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/infra/tracing"
	"google.golang.org/grpc"
)

func LoggingInterceptor(log domain.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = tracing.AddRequestIdGrpc(ctx)
		log.Info(ctx).
			Str("method", info.FullMethod).
			Interface("req", req).
			Msg("starting gRPC call")

		result, err := handler(ctx, req)

		if err == nil {
			log.Info(ctx).
				Str("method", info.FullMethod).
				Interface("result", result).
				Msg("finished gRPC call")
		} else {
			log.Error(ctx).
				Str("method", info.FullMethod).
				Err(err).
				Msg("error gRPC call")
		}

		return result, err
	}
}
