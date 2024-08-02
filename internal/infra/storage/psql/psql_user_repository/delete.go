package psqlUserRepository

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/pkg/etrace"
)

func (u *userRepository) DeleteById(ctx context.Context, id domain.Id) error {
	db := u.getQuery(ctx)

	_, err := db.Exec(ctx, DeleteByIdQuery, id)
	if err != nil {
		tr := traceUserRepository.OfName("DeleteById").
			OfCauseMethod("db.Exec").
			OfCauseParams(etrace.FuncParams{
				"id": id,
			})
		return eerror.
			Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}

	return nil
}
