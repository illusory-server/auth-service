package interceptors

import (
	"context"
	"github.com/illusory-server/auth-service/internal/infra/metrics"
	"google.golang.org/grpc"
	"time"
)

func RequestMetricsInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start).Seconds()
		metrics.RequestCount.WithLabelValues(info.FullMethod).Inc()
		metrics.RequestDuration.WithLabelValues(info.FullMethod).Observe(elapsed)
	}()
	return handler(ctx, req)
}
