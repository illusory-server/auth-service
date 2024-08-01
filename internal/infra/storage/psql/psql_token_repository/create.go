package psqlTokenRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/eerr"
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
		u.log.Error(ctx).
			Err(err).
			Str("id", string(id)).
			Str("token", token).
			Msg("failed to create token")

		return nil, eerr.Wrap(err, "[tokenRepository] db.Exec")
	}

	return resToken, nil
}
