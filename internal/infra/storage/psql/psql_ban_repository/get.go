package psqlBanRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/pkg/eerr"
)

func (b *banRepository) GetBanReasonById(ctx context.Context, id domain.Id) (string, error) {
	db := b.getQuery(ctx)

	reason := ""
	err := db.QueryRow(ctx, GetBanReasonByIdQuery, id).Scan(&reason)
	if err != nil {
		b.log.Error(ctx).Err(err).Str("id", string(id)).Msg("cannot get ban reason")

		return reason, eerr.Wrap(err, "[banRepository.GetBanReasonById] db.QueryRow")
	}

	return reason, nil
}
