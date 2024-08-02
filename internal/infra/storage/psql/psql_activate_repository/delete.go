package psqlActivateRepository

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/pkg/etrace"
)

func (a *activateRepository) DeleteById(ctx context.Context, userId domain.Id) error {
	db := a.getQuery(ctx)

	_, err := db.Exec(ctx, DeleteByIdQuery, userId)
	if err != nil {
		tr := traceActivateRepository.OfName("DeleteById").
			OfCauseMethod("db.Exec").
			OfCauseParams(etrace.FuncParams{
				"id": userId,
			})
		return eerror.Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}

	return nil
}
