package psqlTokenRepository

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/pkg/etrace"
)

func (u *tokenRepository) HasById(ctx context.Context, id domain.Id) (bool, error) {
	db := u.getQuery(ctx)

	var exists bool
	err := db.QueryRow(ctx, HasByIdQuery, id).Scan(&exists)
	if err != nil {
		tr := traceTokenRepository.OfName("HasById").
			OfCauseMethod("db.QueryRow").
			OfCauseParams(etrace.FuncParams{
				"id": id,
			})
		return false, eerror.Err(err).
			Code(eerror.ErrInternal).
			Stack(tr).
			Msg(eerror.MsgInternal).
			Err()
	}

	return exists, err
}

func (u *tokenRepository) HasByValue(ctx context.Context, token string) (bool, error) {
	db := u.getQuery(ctx)

	var exists bool
	err := db.QueryRow(ctx, HasByValueQuery, token).Scan(&exists)
	if err != nil {
		tr := traceTokenRepository.OfName("HasByValue").
			OfCauseMethod("db.QueryRow").
			OfCauseParams(etrace.FuncParams{
				"token": token,
			})
		return false, eerror.Err(err).
			Code(eerror.ErrInternal).
			Stack(tr).
			Msg(eerror.MsgInternal).
			Err()
	}

	return exists, err
}
