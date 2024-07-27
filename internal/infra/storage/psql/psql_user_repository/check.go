package psqlUserRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/pkg/eerr"
)

func (u *userRepository) HasById(ctx context.Context, id domain.Id) (bool, error) {
	db := u.getQuery(ctx)
	var exists bool

	err := db.QueryRow(ctx, HasByIdQuery, id).Scan(&exists)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("id", string(id)).
			Msg("has by id query failure")

		return false, eerr.Wrap(err, "[userRepository.HasById] db.QueryRow")
	}

	return exists, nil
}

func (u *userRepository) HasByLogin(ctx context.Context, login string) (bool, error) {
	db := u.getQuery(ctx)
	var exists bool

	err := db.QueryRow(ctx, HasByLoginQuery, login).Scan(&exists)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("login", login).
			Msg("has by login query failure")

		return false, eerr.Wrap(err, "[userRepository.HasByLogin] db.QueryRow")
	}

	return exists, nil
}

func (u *userRepository) HasByEmail(ctx context.Context, email string) (bool, error) {
	db := u.getQuery(ctx)
	var exists bool

	err := db.QueryRow(ctx, HasByEmailQuery, email).Scan(&exists)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("email", email).
			Msg("has by email query failure")

		return false, eerr.Wrap(err, "[userRepository.HasByEmail] db.QueryRow")
	}

	return exists, nil
}

func (u *userRepository) CheckUserRole(ctx context.Context, id, role string) (bool, error) {
	db := u.getQuery(ctx)
	var check bool

	err := db.QueryRow(ctx, CheckUserRoleQuery, id, role).Scan(&check)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("id", id).
			Str("role", role).
			Msg("check user role failure")

		return false, eerr.Wrap(err, "[userRepository.CheckUserRole] db.QueryRow")
	}

	return check, nil
}
