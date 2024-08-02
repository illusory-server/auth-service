package psqlTokenRepository

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/etrace"
)

func (u *tokenRepository) Save(ctx context.Context, id domain.Id, token string) (*model.Token, error) {
	has, err := u.HasById(ctx, id)
	if err != nil {
		return nil, err
	}
	if !has {
		res, err := u.Create(ctx, id, token)
		if err != nil {
			tr := traceTokenRepository.OfName("Save").
				OfCauseMethod("Create").
				OfCauseParams(etrace.FuncParams{
					"id":    id,
					"token": token,
				})
			return nil, eerror.Err(err).
				Stack(tr).
				Err()
		}
		return res, nil
	}
	res, err := u.UpdateById(ctx, id, token)
	if err != nil {
		tr := traceTokenRepository.OfName("Save").
			OfCauseMethod("UpdateById").
			OfCauseParams(etrace.FuncParams{
				"id":    id,
				"token": token,
			})
		return nil, eerror.Err(err).
			Stack(tr).
			Err()
	}
	return res, nil
}
