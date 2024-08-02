package psqlBanRepository

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/etrace"
	"time"
)

func (b *banRepository) BanById(ctx context.Context, id domain.Id, reason string) (*model.Ban, error) {
	db := b.getQuery(ctx)

	ban := &model.Ban{}
	err := db.QueryRow(ctx, BanByIdQuery, id, reason, time.Now()).Scan(
		&ban.Id,
		&ban.IsBanned,
		&ban.BanReason,
		&ban.UpdatedAt,
		&ban.CreatedAt,
	)
	if err != nil {
		tr := traceBanRepository.OfName("BanById").
			OfCauseMethod("db.QueryRow").
			OfCauseParams(etrace.FuncParams{
				"id":     id,
				"reason": reason,
			})
		return nil, eerror.Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}

	return ban, nil
}

func (b *banRepository) UnbanById(ctx context.Context, id domain.Id) (*model.Ban, error) {
	db := b.getQuery(ctx)

	ban := &model.Ban{}
	err := db.QueryRow(ctx, UnbanByIdQuery, id, time.Now()).Scan(
		&ban.Id,
		&ban.IsBanned,
		&ban.BanReason,
		&ban.UpdatedAt,
		&ban.CreatedAt,
	)
	if err != nil {
		tr := traceBanRepository.OfName("UnbanById").
			OfCauseMethod("db.QueryRow").
			OfCauseParams(etrace.FuncParams{
				"id": id,
			})
		return nil, eerror.Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}

	return ban, nil
}
