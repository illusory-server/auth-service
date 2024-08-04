package grpcv1AuthService

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	authv1 "github.com/illusory-server/auth-service/gen/auth"
	"github.com/illusory-server/auth-service/pkg/etrace"
)

func (a *AuthServiceServer) Refresh(ctx context.Context, token *authv1.RefreshToken) (*authv1.AuthResponse, error) {
	authRes, err := a.authUseCase.Refresh(ctx, token.RefreshToken)
	if err != nil {
		tr := traceAuthService.OfName("Refresh").
			OfCauseMethod("authUseCase.Refresh").
			OfCauseParams(etrace.FuncParams{
				"token": token.RefreshToken,
			})
		return nil, eerror.Err(err).
			Stack(tr).
			Err()
	}
	return authMapper.AuthResultToAuthResponse(authRes), nil
}
