package psqlUserRepository

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/etrace"
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
		tr := traceUserRepository.OfName("Create").
			OfCauseMethod("db.Exec").
			OfParams(etrace.FuncParams{
				"user": user,
			})
		return nil, eerror.
			Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}

	return user, nil
}
