package psqlTokenRepository

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/pkg/etrace"
)

func (u *tokenRepository) DeleteByValue(ctx context.Context, value string) error {
	db := u.getQuery(ctx)

	_, err := db.Exec(ctx, DeleteByValueQuery, value)
	if err != nil {
		tr := traceTokenRepository.OfName("DeleteByValue").
			OfCauseMethod("db.Exec").
			OfCauseParams(etrace.FuncParams{
				"value": value,
			})
		return eerror.Err(err).
			Code(eerror.ErrInternal).
			Stack(tr).
			Msg(eerror.MsgInternal).
			Err()
	}

	return nil
}

func (u *tokenRepository) DeleteById(ctx context.Context, id domain.Id) error {
	db := u.getQuery(ctx)

	_, err := db.Exec(ctx, DeleteByIdQuery, id)
	if err != nil {
		tr := traceTokenRepository.OfName("DeleteById").
			OfCauseMethod("db.Exec").
			OfCauseParams(etrace.FuncParams{
				"id": id,
			})
		return eerror.Err(err).
			Code(eerror.ErrInternal).
			Stack(tr).
			Msg(eerror.MsgInternal).
			Err()
	}

	return nil
}
