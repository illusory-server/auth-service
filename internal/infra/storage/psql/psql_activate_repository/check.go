package psqlActivateRepository

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/pkg/etrace"
)

func (a *activateRepository) IsActivateById(ctx context.Context, userId domain.Id) (bool, error) {
	db := a.getQuery(ctx)

	var exists bool
	err := db.QueryRow(ctx, IsActivateByIdQuery, userId).Scan(&exists)
	if err != nil {
		tr := traceActivateRepository.OfName("IsActivateById").
			OfCauseMethod("db.QueryRow").
			OfCauseParams(etrace.FuncParams{
				"id": userId,
			})
		return false, eerror.Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}

	return exists, nil
}

func (a *activateRepository) HasById(ctx context.Context, userId domain.Id) (bool, error) {
	db := a.getQuery(ctx)

	var exists bool
	err := db.QueryRow(ctx, HasByIdQuery, userId).Scan(&exists)
	if err != nil {
		tr := traceActivateRepository.OfName("HasById").
			OfCauseMethod("db.QueryRow").
			OfCauseParams(etrace.FuncParams{
				"id": userId,
			})
		return false, eerror.Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}

	return exists, nil
}
