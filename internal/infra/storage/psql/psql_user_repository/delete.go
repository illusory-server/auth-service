package psqlUserRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/pkg/eerr"
)

func (u *userRepository) DeleteById(ctx context.Context, id domain.Id) error {
	db := u.getQuery(ctx)

	_, err := db.Exec(ctx, DeleteByIdQuery, id)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("id", "id").
			Msg("user delete error")

		return eerr.Wrap(err, "[userRepository] db.Exec")
	}

	return nil
}
