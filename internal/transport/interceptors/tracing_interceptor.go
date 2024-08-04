package interceptors

import (
	"context"
	"github.com/illusory-server/auth-service/internal/infra/tracing"
	"google.golang.org/grpc"
)

func Tracing(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	ctx = tracing.AddRequestIdGrpc(ctx)
	return handler(ctx, req)
}
