package psql

import (
	"context"
	"fmt"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/sql"
	"github.com/illusory-server/auth-service/internal/infra/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func Connect(config *config.Config, log domain.Logger) sql.QueryExecutor {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Postgres.Host, config.Postgres.Port, config.Postgres.User, config.Postgres.Password, config.Postgres.DbName)

	cfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		panic(err)
	}

	cfg.MaxConns = 20
	cfg.MinConns = 4
	cfg.MaxConnLifetime = time.Hour
	cfg.MaxConnIdleTime = time.Minute * 20
	cfg.HealthCheckPeriod = time.Minute * 2

	poolConn, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		panic(err)
	}

	return poolConn
}
