package psqlActivateRepository

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/etrace"
	"time"
)

func (a *activateRepository) Update(ctx context.Context, activate *model.Activate) (*model.Activate, error) {
	db := a.getQuery(ctx)

	now := time.Now()
	_, err := db.Exec(ctx, UpdateQuery, activate.Id, activate.IsActivated, activate.Link, now)
	if err != nil {
		tr := traceActivateRepository.OfName("Update").
			OfCauseMethod("db.Exec").
			OfCauseParams(etrace.FuncParams{
				"update_activate": activate,
			})
		return nil, eerror.Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}
	activate.UpdatedAt = now
	return activate, nil
}

func (a *activateRepository) ActivateUserById(ctx context.Context, userId domain.Id) error {
	db := a.getQuery(ctx)

	_, err := db.Exec(ctx, ActivateUserByIdQuery, userId, time.Now())
	if err != nil {
		tr := traceActivateRepository.OfName("ActivateUserById").
			OfCauseMethod("db.Exec").
			OfCauseParams(etrace.FuncParams{
				"id": userId,
			})
		return eerror.Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}
	return nil
}

func (a *activateRepository) ActivateUserByLink(ctx context.Context, link string) error {
	db := a.getQuery(ctx)

	_, err := db.Exec(ctx, ActivateUserByLinkQuery, link, time.Now())
	if err != nil {
		tr := traceActivateRepository.OfName("ActivateUserByLink").
			OfCauseMethod("db.Exec").
			OfCauseParams(etrace.FuncParams{
				"link": link,
			})
		return eerror.Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}
	return nil
}
