package psqlTokenRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/pkg/eerr"
)

func (u *tokenRepository) HasById(ctx context.Context, id domain.Id) (bool, error) {
	db := u.getQuery(ctx)

	var exists bool
	err := db.QueryRow(ctx, HasByIdQuery, id).Scan(&exists)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("id", string(id)).
			Msg("check token error")

		return false, eerr.Wrap(err, "[tokenRepository.HasById] db.QueryRow")
	}

	return exists, err
}

func (u *tokenRepository) HasByValue(ctx context.Context, token string) (bool, error) {
	db := u.getQuery(ctx)

	var exists bool
	err := db.QueryRow(ctx, HasByValueQuery, token).Scan(&exists)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("token", token).
			Msg("check token error")

		return false, eerr.Wrap(err, "[tokenRepository.HasByValue] db.QueryRow")
	}

	return exists, err
}
