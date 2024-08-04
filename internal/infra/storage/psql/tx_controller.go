package psql

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/sql"
	"github.com/illusory-server/auth-service/pkg/etrace"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionKey string

const (
	PgxTransactionKey TransactionKey = "pgx_transaction_key"
)

var (
	traceTransaction = etrace.Method{
		Package: "psql",
		Type:    "PgxTransaction",
		Name:    "WithinTransaction",
	}
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
		p.Log.Debug(ctx).Msg("starting psql transaction failed")
		tr := traceTransaction.OfCauseMethod("Conn.Begin")
		return eerror.Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}

	err = callback(p.TxController.InjectTransaction(ctx, tx))
	if err != nil {
		errLoc := tx.Rollback(ctx)
		if errLoc != nil {
			p.Log.Debug(ctx).Msg("rolling back psql transaction failed")
			tr := traceTransaction.OfCauseMethod("tx.Rollback(1)")
			return eerror.Err(errLoc).
				Code(eerror.ErrInternal).
				Msg(eerror.MsgInternal).
				Stack(tr).
				Err()
		}
		tr := traceTransaction.OfCauseMethod("callback")
		p.Log.Debug(ctx).Msg("rolling back psql transaction")
		return eerror.Err(err).
			Stack(tr).
			Err()
	}

	err = tx.Commit(ctx)
	if err != nil {
		errLoc := tx.Rollback(ctx)
		if errLoc != nil {
			p.Log.Debug(ctx).Msg("rolling back psql transaction failed")
			tr := traceTransaction.OfCauseMethod("tx.Rollback(2)")
			return eerror.Err(errLoc).
				Code(eerror.ErrInternal).
				Msg(eerror.MsgInternal).
				Stack(tr).
				Err()
		}
		p.Log.Debug(ctx).Msg("commit psql transaction failed")
		tr := traceTransaction.OfCauseMethod("tx.Commit")
		return eerror.Err(err).
			Stack(tr).
			Err()
	}
	p.Log.Debug(ctx).Msg("commit psql transaction succeeded")
	return nil
}
