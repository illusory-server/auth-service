package psqlBanRepository

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/pkg/etrace"
)

func (b *banRepository) HasById(ctx context.Context, id domain.Id) (bool, error) {
	db := b.getQuery(ctx)

	var exists bool
	err := db.QueryRow(ctx, HasByIdQuery, id).Scan(&exists)
	if err != nil {
		tr := traceBanRepository.OfName("HasById").
			OfCauseMethod("db.QueryRow").
			OfCauseParams(etrace.FuncParams{
				"id": id,
			})
		return false, eerror.Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}

	return exists, nil
}

func (b *banRepository) IsBannedById(ctx context.Context, id domain.Id) (bool, error) {
	db := b.getQuery(ctx)

	var exists bool
	err := db.QueryRow(ctx, IsBannedByIdQuery, id).Scan(&exists)
	if err != nil {
		tr := traceBanRepository.OfName("IsBannedById").
			OfCauseMethod("db.QueryRow").
			OfCauseParams(etrace.FuncParams{
				"id": id,
			})
		return false, eerror.Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}

	return exists, nil
}
