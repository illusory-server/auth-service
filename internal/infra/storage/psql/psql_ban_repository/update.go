package psqlBanRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/eerr"
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
		b.log.Error(ctx).Err(err).Str("id", string(id)).Str("reason", reason).Msg("cannot ban user")

		return nil, eerr.Wrap(err, "[banRepository.BanById] db.QueryRow")
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
		b.log.Error(ctx).Err(err).Str("id", string(id)).Msg("cannot unban user")

		return nil, eerr.Wrap(err, "[banRepository.UnbanById] db.QueryRow")
	}

	return ban, nil
}
