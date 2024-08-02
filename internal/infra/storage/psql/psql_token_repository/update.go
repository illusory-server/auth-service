package psqlTokenRepository

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/etrace"
	"time"
)

func (u *tokenRepository) UpdateById(ctx context.Context, id domain.Id, token string) (*model.Token, error) {
	db := u.getQuery(ctx)

	resToken := &model.Token{}
	err := db.QueryRow(ctx, UpdateQuery, id, token, time.Now()).Scan(&resToken.Id, &resToken.Value, &resToken.UpdatedAt, &resToken.CreatedAt)
	if err != nil {
		tr := traceTokenRepository.OfName("UpdateById").
			OfCauseMethod("db.QueryRow").
			OfCauseParams(etrace.FuncParams{
				"id":    id,
				"token": token,
			})
		return nil, eerror.Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}

	return resToken, nil
}
