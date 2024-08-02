package psqlActivateRepository

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/etrace"
)

func (a *activateRepository) GetByUserId(ctx context.Context, userId domain.Id) (*model.Activate, error) {
	db := a.getQuery(ctx)

	activate := &model.Activate{}
	err := db.QueryRow(ctx, GetByUserIdQuery, userId).Scan(
		&activate.Id,
		&activate.IsActivated,
		&activate.Link,
		&activate.UpdatedAt,
		&activate.CreatedAt,
	)
	if err != nil {
		tr := traceActivateRepository.OfName("GetByUserId").
			OfCauseMethod("db.Exec").
			OfCauseParams(etrace.FuncParams{
				"id": userId,
			})
		return nil, eerror.Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}

	return activate, nil
}
