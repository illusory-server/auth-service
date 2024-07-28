package psql

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/pkg/eerr"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func PingUnaryInterceptor(pool *pgxpool.Pool, log domain.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Debug(ctx).Msg("ping interceptor work")
		err := pool.Ping(ctx)
		if err != nil {
			log.Error(ctx).Msg("ping to pg failed")
			return nil, status.Error(codes.Internal, eerr.MsgInternal)
		}
		return handler(ctx, req)
	}
}
