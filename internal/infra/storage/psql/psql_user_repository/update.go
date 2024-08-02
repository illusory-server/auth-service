package psqlUserRepository

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/etrace"
	"time"
)

func (u *userRepository) UpdateById(ctx context.Context, user *model.User) (*model.User, error) {
	db := u.getQuery(ctx)

	_, err := db.Exec(ctx, UpdateByIdQuery, user.Id, user.Login, user.Email, user.Role, time.Now())
	if err != nil {
		return nil, eerror.
			Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(traceUserRepository.
				OfName("UpdateById").
				OfCauseMethod("db.Exec"),
			).
			Err()
	}
	return user, nil
}

func (u *userRepository) UpdateRoleById(ctx context.Context, id domain.Id, role string) (*model.User, error) {
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
		return nil, eerror.Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(traceUserRepository.
				OfName("UpdateRoleById").
				OfCauseMethod("db.Exec").
				OfCauseParams(etrace.FuncParams{
					"id":   id,
					"role": role,
				}),
			).
			Err()
	}
	return user, nil
}

func (u *userRepository) UpdatePasswordById(ctx context.Context, id domain.Id, password string) (*model.User, error) {
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
		return nil, eerror.
			Err(err).
			Code(eerror.ErrInternal).
			Stack(
				traceUserRepository.
					OfName("UpdatePasswordById").
					OfCauseMethod("db.Exec").
					OfCauseParams(etrace.FuncParams{
						"id":       id,
						"password": password,
					}),
			).
			Msg(eerror.MsgInternal).
			Err()
	}
	return user, nil
}
