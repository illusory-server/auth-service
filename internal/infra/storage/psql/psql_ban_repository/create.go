package psqlBanRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/eerr"
)

func (b *banRepository) Create(ctx context.Context, ban *model.Ban) (*model.Ban, error) {
	db := b.getQuery(ctx)

	_, err := db.Exec(ctx, CreateQuery, ban.Id, ban.IsBanned, ban.BanReason, ban.UpdatedAt, ban.CreatedAt)
	if err != nil {
		b.log.Error(ctx).
			Err(err).
			Interface("create_data", ban).
			Msg("cannot create ban")

		return nil, eerr.Wrap(err, "[banRepository.Create] db.Exec")
	}

	return ban, nil
}
