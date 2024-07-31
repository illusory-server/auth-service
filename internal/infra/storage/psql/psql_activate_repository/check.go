package psqlActivateRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/pkg/eerr"
)

func (a *activateRepository) IsActivateById(ctx context.Context, userId domain.Id) (bool, error) {
	db := a.getQuery(ctx)

	var exists bool
	err := db.QueryRow(ctx, IsActivateByIdQuery, userId).Scan(&exists)
	if err != nil {
		a.log.Error(ctx).
			Err(err).
			Str("id", string(userId)).
			Msg("error is activate by id")

		return false, eerr.Wrap(err, "[activateRepository.IsActivateById] db.QueryRow")
	}

	return exists, nil
}

func (a *activateRepository) HasById(ctx context.Context, userId domain.Id) (bool, error) {
	db := a.getQuery(ctx)

	var exists bool
	err := db.QueryRow(ctx, HasByIdQuery, userId).Scan(&exists)
	if err != nil {
		a.log.Error(ctx).
			Err(err).
			Str("id", string(userId)).
			Msg("error has by id")

		return false, eerr.Wrap(err, "[activateRepository.HasById] db.QueryRow")
	}

	return exists, nil
}
