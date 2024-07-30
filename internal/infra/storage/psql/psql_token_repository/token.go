package psqlTokenRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/repository"
	"github.com/illusory-server/auth-service/internal/domain/sql"
)

type tokenRepository struct {
	db  sql.QueryExecutor
	log domain.Logger
	tx  sql.TransactionController
}

func (u *tokenRepository) getQuery(ctx context.Context) sql.QueryExecutor {
	db := u.tx.ExtractTransaction(ctx)
	if db != nil {
		return db
	}
	return u.db
}

func New(db sql.QueryExecutor, logger domain.Logger, tx sql.TransactionController) repository.TokenRepository {
	return &tokenRepository{
		db:  db,
		log: logger,
		tx:  tx,
	}
}
