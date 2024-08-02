package psqlUserRepository

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/pkg/etrace"
)

func (u *userRepository) HasById(ctx context.Context, id domain.Id) (bool, error) {
	db := u.getQuery(ctx)
	var exists bool

	err := db.QueryRow(ctx, HasByIdQuery, id).Scan(&exists)
	if err != nil {
		tr := traceUserRepository.OfName("HasById").
			OfCauseMethod("db.QueryRow").
			OfCauseParams(etrace.FuncParams{
				"id": id,
			})
		return false, eerror.
			Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}

	return exists, nil
}

func (u *userRepository) HasByLogin(ctx context.Context, login string) (bool, error) {
	db := u.getQuery(ctx)
	var exists bool

	err := db.QueryRow(ctx, HasByLoginQuery, login).Scan(&exists)
	if err != nil {
		tr := traceUserRepository.OfName("HasByLogin").
			OfCauseMethod("db.QueryRow").
			OfCauseParams(etrace.FuncParams{
				"login": login,
			})
		return false, eerror.
			Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}

	return exists, nil
}

func (u *userRepository) HasByEmail(ctx context.Context, email string) (bool, error) {
	db := u.getQuery(ctx)
	var exists bool

	err := db.QueryRow(ctx, HasByEmailQuery, email).Scan(&exists)
	if err != nil {
		tr := traceUserRepository.OfName("HasByEmail").
			OfCauseMethod("db.QueryRow").
			OfCauseParams(etrace.FuncParams{
				"email": email,
			})
		return false, eerror.
			Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}

	return exists, nil
}

func (u *userRepository) CheckUserRole(ctx context.Context, id domain.Id, role string) (bool, error) {
	db := u.getQuery(ctx)
	var check bool

	err := db.QueryRow(ctx, CheckUserRoleQuery, id, role).Scan(&check)
	if err != nil {
		tr := traceUserRepository.OfName("CheckUserRole").
			OfCauseMethod("db.QueryRow").
			OfCauseParams(etrace.FuncParams{
				"id":   id,
				"role": role,
			})
		return false, eerror.
			Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}

	return check, nil
}
