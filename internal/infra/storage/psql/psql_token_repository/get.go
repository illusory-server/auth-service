package psqlTokenRepository

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/etrace"
)

func (u *tokenRepository) GetById(ctx context.Context, id domain.Id) (*model.Token, error) {
	db := u.getQuery(ctx)

	token := &model.Token{}
	err := db.QueryRow(ctx, GetByIdQuery, id).Scan(
		&token.Id,
		&token.Value,
		&token.UpdatedAt,
		&token.CreatedAt,
	)

	if err != nil {
		tr := traceTokenRepository.OfName("GetById").
			OfCauseMethod("db.QueryRow").
			OfCauseParams(etrace.FuncParams{
				"id": id,
			})
		return nil, eerror.Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}

	return token, nil
}
