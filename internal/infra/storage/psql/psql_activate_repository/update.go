package psqlActivateRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/eerr"
	"time"
)

func (a *activateRepository) Update(ctx context.Context, activate *model.Activate) (*model.Activate, error) {
	db := a.getQuery(ctx)

	now := time.Now()
	_, err := db.Exec(ctx, UpdateQuery, activate.Id, activate.IsActivated, activate.Link, now)
	if err != nil {
		a.log.Error(ctx).
			Err(err).
			Interface("update_data", activate).
			Msg("failed to update activate")

		return nil, eerr.Wrap(err, "[activateRepository.Update] db.Exec")
	}
	activate.UpdatedAt = now
	return activate, nil
}

func (a *activateRepository) ActivateUserById(ctx context.Context, userId domain.Id) error {
	db := a.getQuery(ctx)

	_, err := db.Exec(ctx, ActivateUserByIdQuery, userId, time.Now())
	if err != nil {
		a.log.Error(ctx).
			Err(err).
			Str("id", string(userId)).
			Msg("failed to activate user by id")

		return eerr.Wrap(err, "[activateRepository.ActivateUserById] db.Exec")
	}
	return nil
}

func (a *activateRepository) ActivateUserByLink(ctx context.Context, link string) error {
	db := a.getQuery(ctx)

	_, err := db.Exec(ctx, ActivateUserByLinkQuery, link, time.Now())
	if err != nil {
		a.log.Error(ctx).
			Err(err).
			Str("link", link).
			Msg("failed to activate user by link")

		return eerr.Wrap(err, "[activateRepository.ActivateUserById] db.Exec")
	}
	return nil
}
