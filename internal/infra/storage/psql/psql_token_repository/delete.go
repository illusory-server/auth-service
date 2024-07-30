package psqlTokenRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/pkg/eerr"
)

func (u *tokenRepository) DeleteByValue(ctx context.Context, value string) error {
	db := u.getQuery(ctx)

	_, err := db.Exec(ctx, DeleteByValueQuery, value)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("value", value).
			Msg("failed to delete token by value")

		return eerr.Wrap(err, "[tokenRepository.DeleteByValue]: db.Exec")
	}

	return nil
}

func (u *tokenRepository) DeleteById(ctx context.Context, id domain.Id) error {
	db := u.getQuery(ctx)

	_, err := db.Exec(ctx, DeleteByIdQuery, id)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("id", string(id)).
			Msg("failed to delete token by value")

		return eerr.Wrap(err, "[tokenRepository.DeleteByValue]: db.Exec")
	}

	return nil
}
