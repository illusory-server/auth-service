package psqlActivateRepository

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/etrace"
	"time"
)

func (a *activateRepository) Create(ctx context.Context, userId domain.Id, link string) (*model.Activate, error) {
	db := a.getQuery(ctx)

	now := time.Now()
	activate := &model.Activate{
		Id:          userId,
		IsActivated: false,
		Link:        link,
		UpdatedAt:   now,
		CreatedAt:   now,
	}
	_, err := db.Exec(ctx, CreateQuery, userId, activate.IsActivated, activate.Link, activate.UpdatedAt, activate.CreatedAt)
	if err != nil {
		tr := traceActivateRepository.OfName("Create").
			OfCauseMethod("db.Exec").
			OfCauseParams(etrace.FuncParams{
				"id":   userId,
				"link": link,
			})
		return nil, eerror.Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}

	return activate, nil
}
