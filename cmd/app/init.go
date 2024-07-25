package app

import (
	"github.com/illusory-server/auth-service/internal/infra/config"
	"github.com/illusory-server/auth-service/internal/infra/logger"
	"github.com/illusory-server/auth-service/internal/infra/storage/psql"
)

func Init() *App {
	cfg := config.MustLoad()
	log := logger.MustLoad(cfg.Env)
	psqlConn := psql.Connect(cfg, log)
	application := &App{
		Logger: log,
		Cfg:    cfg,
		PSQL:   psqlConn,
	}
	return application
}
