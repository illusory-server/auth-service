package authUseCase

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/pkg/etrace"
)

func (u *useCase) Logout(ctx context.Context, refreshToken string) error {
	err := u.tokenRepository.DeleteByValue(ctx, refreshToken)
	if err != nil {
		tr := traceUseCase.OfName("Logout").
			OfCauseMethod("tokenRepository.DeleteByValue").
			OfCauseParams(etrace.FuncParams{
				"refresh_token": refreshToken,
			})
		return eerror.Err(err).Stack(tr).Err()
	}
	return nil
}
