package psqlActivateRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/repository"
	"github.com/illusory-server/auth-service/internal/domain/sql"
)

type activateRepository struct {
	log domain.Logger
	tx  sql.TransactionController
	db  sql.QueryExecutor
}

func (a *activateRepository) getQuery(ctx context.Context) sql.QueryExecutor {
	db := a.tx.ExtractTransaction(ctx)
	if db != nil {
		return db
	}
	return a.db
}

func New(log domain.Logger, tx sql.TransactionController, db sql.QueryExecutor) repository.ActivateRepository {
	return &activateRepository{
		log: log,
		tx:  tx,
		db:  db,
	}
}
