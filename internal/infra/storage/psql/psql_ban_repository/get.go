package psqlBanRepository

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/pkg/etrace"
)

func (b *banRepository) GetBanReasonById(ctx context.Context, id domain.Id) (string, error) {
	db := b.getQuery(ctx)

	reason := ""
	err := db.QueryRow(ctx, GetBanReasonByIdQuery, id).Scan(&reason)
	if err != nil {
		tr := traceBanRepository.OfName("GetBanReasonById").
			OfCauseMethod("db.QueryRow").
			OfCauseParams(etrace.FuncParams{
				"id": id,
			})
		return reason, eerror.Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}

	return reason, nil
}
