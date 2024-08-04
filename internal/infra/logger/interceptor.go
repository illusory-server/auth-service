package logger

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"google.golang.org/grpc"
)

func LoggingInterceptor(log domain.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
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
		}

		return result, err
	}
}
