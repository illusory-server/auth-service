package psqlUserRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/eerr"
	"time"
)

func (u *userRepository) UpdateById(ctx context.Context, user *model.User) (*model.User, error) {
	db := u.getQuery(ctx)

	_, err := db.Exec(ctx, UpdateByIdQuery, user.Id, user.Login, user.Email, user.Role, time.Now())
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Interface("user", user).
			Msg("update user failed")

		return nil, eerr.Wrap(err, "[userRepository.UpdateById] db.Exec")
	}
	return user, nil
}

func (u *userRepository) UpdateRoleById(ctx context.Context, id string, role string) (*model.User, error) {
	db := u.getQuery(ctx)
	user := &model.User{}
	err := db.QueryRow(ctx, UpdateRoleByIdQuery, id, role, time.Now()).Scan(
		&user.Id,
		&user.Login,
		&user.Email,
		&user.Role,
		&user.Password,
		&user.UpdatedAt,
		&user.CreatedAt,
	)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("id", id).
			Str("role", role).
			Msg("update user failed")

		return nil, eerr.Wrap(err, "[userRepository.UpdateRoleById] db.QueryRow")
	}
	return user, nil
}

func (u *userRepository) UpdatePasswordById(ctx context.Context, id string, password string) (*model.User, error) {
	db := u.getQuery(ctx)
	user := &model.User{}
	err := db.QueryRow(ctx, UpdatePasswordByIdQuery, id, password, time.Now()).Scan(
		&user.Id,
		&user.Login,
		&user.Email,
		&user.Role,
		&user.Password,
		&user.UpdatedAt,
		&user.CreatedAt,
	)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("id", id).
			Str("role", password).
			Msg("update user failed")

		return nil, eerr.Wrap(err, "[userRepository.UpdatePasswordById] db.QueryRow")
	}
	return user, nil
}
