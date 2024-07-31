package psqlActivateRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/eerr"
)

func (a *activateRepository) GetByUserId(ctx context.Context, userId domain.Id) (*model.Activate, error) {
	db := a.getQuery(ctx)

	activate := &model.Activate{}
	err := db.QueryRow(ctx, GetByUserIdQuery, userId).Scan(
		&activate.Id,
		&activate.IsActivated,
		&activate.Link,
		&activate.UpdatedAt,
		&activate.CreatedAt,
	)
	if err != nil {
		a.log.Error(ctx).
			Err(err).
			Str("userId", string(userId)).
			Msg("cannot get activate by id")

		return nil, eerr.Wrap(err, "[activateRepository.GetByUserId] db.QueryRow")
	}

	return activate, nil
}
