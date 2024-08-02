package psqlTokenRepository

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/etrace"
	"time"
)

func (u *tokenRepository) Create(ctx context.Context, id domain.Id, token string) (*model.Token, error) {
	db := u.getQuery(ctx)

	resToken := &model.Token{
		Id:        id,
		Value:     token,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := db.Exec(ctx, CreateQuery, id, token, resToken.UpdatedAt, resToken.CreatedAt)
	if err != nil {
		tr := traceTokenRepository.OfName("Create").
			OfCauseMethod("db.Exec").
			OfCauseParams(etrace.FuncParams{
				"id":    id,
				"token": token,
			})
		return nil, eerror.Err(err).
			Code(eerror.ErrInternal).
			Stack(tr).
			Msg(eerror.MsgInternal).
			Err()
	}

	return resToken, nil
}
