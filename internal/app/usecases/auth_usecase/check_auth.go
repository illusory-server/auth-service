package authUseCase

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	tokenService "github.com/illusory-server/auth-service/internal/app/services/token_service"
	"github.com/illusory-server/auth-service/pkg/etrace"
)

func (u *useCase) CheckAuth(ctx context.Context, accessToken string) (*tokenService.JwtUserData, error) {
	result, err := u.tokenService.ValidateAccessToken(ctx, accessToken)
	if err != nil {
		tr := traceUseCase.OfName("CheckAuth").
			OfCauseMethod("tokenService.ValidateAccessToken").
			OfCauseParams(etrace.FuncParams{
				"access_token": accessToken,
			})
		return nil, eerror.Err(err).
			Stack(tr).
			Err()
	}
	return result, nil
}
