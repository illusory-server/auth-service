package psqlBanRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/repository"
	"github.com/illusory-server/auth-service/internal/domain/sql"
)

type banRepository struct {
	log domain.Logger
	tx  sql.TransactionController
	db  sql.QueryExecutor
}

func (b *banRepository) getQuery(ctx context.Context) sql.QueryExecutor {
	db := b.tx.ExtractTransaction(ctx)
	if db != nil {
		return db
	}
	return b.db
}

func New(log domain.Logger, tx sql.TransactionController, db sql.QueryExecutor) repository.BanRepository {
	return &banRepository{
		log: log,
		tx:  tx,
		db:  db,
	}
}
