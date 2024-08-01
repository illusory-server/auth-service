package psql

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/sql"
	"github.com/illusory-server/auth-service/pkg/eerr"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionKey string

const (
	PgxTransactionKey TransactionKey = "pgx_transaction_key"
)

type PgxTransaction struct {
	Conn         *pgxpool.Pool
	TxController sql.TransactionController
	Log          domain.Logger
}

type PgxTransactionController struct{}

func (t PgxTransactionController) InjectTransaction(ctx context.Context, tx sql.QueryExecutor) context.Context {
	return context.WithValue(ctx, PgxTransactionKey, tx)
}

func (t PgxTransactionController) ExtractTransaction(ctx context.Context) sql.QueryExecutor {
	if tx, ok := ctx.Value(PgxTransactionKey).(sql.QueryExecutor); ok {
		return tx
	}
	return nil
}

func (p *PgxTransaction) WithinTransaction(ctx context.Context, callback func(context.Context) error) error {
	tx, err := p.Conn.Begin(ctx)
	if err != nil {
		p.Log.Error(ctx).
			Msg("starting psql transaction failed")

		return eerr.Wrap(err, "[PgxTransaction.WithinTransaction] Conn.Begin")
	}

	err = callback(p.TxController.InjectTransaction(ctx, tx))
	if err != nil {
		errLoc := tx.Rollback(ctx)
		if errLoc != nil {
			p.Log.Error(ctx).
				Msg("rolling back psql transaction failed")
			return eerr.Wrap(err, "[PgxTransaction.WithinTransaction] tx.Rollback(1)")
		}
		p.Log.Info(ctx).
			Msg("rolling back psql transaction")
		return eerr.Wrap(err, "[PgxTransaction.WithinTransaction] callback")
	}

	err = tx.Commit(ctx)
	if err != nil {
		err := tx.Rollback(ctx)
		if err != nil {
			p.Log.Error(ctx).
				Msg("rolling back psql transaction failed")
			return eerr.Wrap(err, "[PgxTransaction.WithinTransaction] tx.Rollback(2)")
		}
		p.Log.Error(ctx).
			Msg("commit psql transaction failed")
		return err
	}
	p.Log.Info(ctx).
		Msg("commit psql transaction succeeded")
	return nil
}
