package psqlTokenRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"time"
)

func (u *tokenRepository) UpdateById(ctx context.Context, id domain.Id, token string) (*model.Token, error) {
	db := u.getQuery(ctx)

	resToken := &model.Token{}
	err := db.QueryRow(ctx, UpdateQuery, id, token, time.Now()).Scan(&resToken.Id, &resToken.Value, &resToken.UpdatedAt, &resToken.CreatedAt)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("id", string(id)).
			Str("token", token).
			Msg("update token error")

		return nil, err
	}

	return resToken, nil
}
