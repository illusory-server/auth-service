package psqlBanRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/pkg/eerr"
)

func (b *banRepository) DeleteById(ctx context.Context, id domain.Id) error {
	db := b.getQuery(ctx)

	_, err := db.Exec(ctx, DeleteByIdQuery, id)
	if err != nil {
		b.log.Error(ctx).Err(err).Str("id", string(id)).Msg("error while deleting ban by id")

		return eerr.Wrap(err, "[banRepository.DeleteById] db.Exec")
	}

	return nil
}
