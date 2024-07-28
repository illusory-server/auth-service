package psqlTokenRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/eerr"
)

func (u *tokenRepository) Save(ctx context.Context, id domain.Id, token string) (*model.Token, error) {
	has, err := u.HasById(ctx, id)
	if err != nil {
		return nil, err
	}
	if !has {
		res, err := u.Create(ctx, id, token)
		if err != nil {
			u.log.Error(ctx).
				Err(err).
				Str("token", token).
				Str("id", string(id)).
				Msg("error creating token")

			return nil, eerr.Wrap(err, "[tokenRepository.Save] Create")
		}
		return res, nil
	}
	res, err := u.UpdateById(ctx, id, token)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("token", token).
			Str("id", string(id)).
			Msg("error updating token")

		return nil, eerr.Wrap(err, "[tokenRepository.Save] UpdateById")
	}
	return res, nil
}
