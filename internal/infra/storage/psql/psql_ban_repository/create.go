package psqlBanRepository

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/etrace"
)

func (b *banRepository) Create(ctx context.Context, ban *model.Ban) (*model.Ban, error) {
	db := b.getQuery(ctx)

	_, err := db.Exec(ctx, CreateQuery, ban.Id, ban.IsBanned, ban.BanReason, ban.UpdatedAt, ban.CreatedAt)
	if err != nil {
		tr := traceBanRepository.OfName("Create").
			OfCauseMethod("db.Exec").
			OfCauseParams(etrace.FuncParams{
				"created_ban": ban,
			})
		return nil, eerror.Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}

	return ban, nil
}
