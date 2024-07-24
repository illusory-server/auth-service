package main

import (
	"context"
	"github.com/illusory-server/auth-service/internal/infra/config"
	"github.com/illusory-server/auth-service/internal/infra/logger"
	"github.com/illusory-server/auth-service/internal/infra/tracing"
)

func main() {
	cfg := config.MustLoad()
	log := logger.MustLoad(cfg.Env)
	ctx := context.Background()
	ctx = tracing.AddRequestId(ctx)
	log.Info(ctx).
		Msg("config and load initialized")
	log.Error(ctx).
		Msg("config and load initialized")
	log.Error(ctx).
		Msg("config and load initialized")
}
