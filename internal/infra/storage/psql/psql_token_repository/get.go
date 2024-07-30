package psqlTokenRepository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/eerr"
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
		u.log.Error(ctx).
			Err(err).
			Str("id", string(id)).
			Msg("get token by id failed")

		return nil, eerr.Wrap(err, "[tokenRepository.GetById]: db.QueryRow")
	}

	return token, nil
}
