package psqlActivateRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/pkg/eerr"
)

func (a *activateRepository) DeleteById(ctx context.Context, userId domain.Id) error {
	db := a.getQuery(ctx)

	_, err := db.Exec(ctx, DeleteByIdQuery, userId)
	if err != nil {
		a.log.Error(ctx).
			Err(err).
			Str("id", string(userId)).
			Msg("error deleting user")

		return eerr.Wrap(err, "[activateRepository.DeleteById] db.Exec")
	}

	return nil
}
