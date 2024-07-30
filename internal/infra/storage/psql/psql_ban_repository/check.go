package psqlBanRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/pkg/eerr"
)

func (b *banRepository) HasById(ctx context.Context, id domain.Id) (bool, error) {
	db := b.getQuery(ctx)

	var exists bool
	err := db.QueryRow(ctx, HasByIdQuery, id).Scan(&exists)
	if err != nil {
		b.log.Error(ctx).
			Err(err).
			Str("id", string(id)).
			Msg("error has ban by id query")

		return false, eerr.Wrap(err, "[banRepository.HasById] db.QueryRow")
	}

	return exists, nil
}

func (b *banRepository) IsBannedById(ctx context.Context, id domain.Id) (bool, error) {
	db := b.getQuery(ctx)

	var exists bool
	err := db.QueryRow(ctx, IsBannedByIdQuery, id).Scan(&exists)
	if err != nil {
		b.log.Error(ctx).
			Err(err).
			Str("id", string(id)).
			Msg("error is baned by id query")

		return false, eerr.Wrap(err, "[banRepository.IsBannedById] db.QueryRow")
	}

	return exists, nil
}
