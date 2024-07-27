package psqlUserRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/eerr"
)

func (u *userRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	db := u.getQuery(ctx)
	_, err := db.Exec(
		ctx,
		CreateQuery,
		user.Id,
		user.Login,
		user.Email,
		user.Role,
		user.Password,
		user.UpdatedAt,
		user.CreatedAt,
	)

	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Interface("user", user).
			Msg("create user failed")

		return nil, eerr.Wrap(err, "[userRepository] db.Exec")
	}

	return user, nil
}
