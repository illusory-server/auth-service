package psqlUserRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/repository"
	"github.com/illusory-server/auth-service/internal/domain/sql"
)

type userRepository struct {
	db  sql.QueryExecutor
	log domain.Logger
	tx  sql.TransactionController
}

func (u *userRepository) getQuery(ctx context.Context) sql.QueryExecutor {
	db := u.tx.ExtractTransaction(ctx)
	if db != nil {
		return db
	}
	return u.db
}

func New(db sql.QueryExecutor, logger domain.Logger, tx sql.TransactionController) repository.UserRepository {
	return &userRepository{
		db:  db,
		log: logger,
		tx:  tx,
	}
}
