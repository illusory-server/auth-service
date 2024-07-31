package psqlActivateRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/eerr"
	"time"
)

func (a *activateRepository) Create(ctx context.Context, userId domain.Id, link string) (*model.Activate, error) {
	db := a.getQuery(ctx)

	now := time.Now()
	activate := &model.Activate{
		Id:          userId,
		IsActivated: false,
		Link:        link,
		UpdatedAt:   now,
		CreatedAt:   now,
	}
	_, err := db.Exec(ctx, CreateQuery, userId, activate.IsActivated, activate.Link, activate.UpdatedAt, activate.CreatedAt)
	if err != nil {
		a.log.Error(ctx).
			Err(err).
			Str("id", string(userId)).
			Str("link", link).
			Msg("error creating activate")

		return nil, eerr.Wrap(err, "[activateRepository] db.Exec")
	}

	return activate, nil
}
