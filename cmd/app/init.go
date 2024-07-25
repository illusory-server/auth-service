package app

import (
	"github.com/illusory-server/auth-service/internal/infra/config"
	"github.com/illusory-server/auth-service/internal/infra/logger"
)

func Init() *App {
	cfg := config.MustLoad()
	log := logger.MustLoad(cfg.Env)
	application := &App{
		Logger: log,
		Cfg:    cfg,
	}
	return application
}
