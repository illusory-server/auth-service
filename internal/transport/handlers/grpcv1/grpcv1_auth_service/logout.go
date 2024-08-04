package grpcv1AuthService

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	authv1 "github.com/illusory-server/auth-service/gen/auth"
	"github.com/illusory-server/auth-service/pkg/etrace"
)

func (a *AuthServiceServer) Logout(ctx context.Context, token *authv1.RefreshToken) (*authv1.Empty, error) {
	err := a.authUseCase.Logout(ctx, token.RefreshToken)
	if err != nil {
		tr := traceAuthService.OfName("Logout").
			OfCauseMethod("authUseCase.Logout").
			OfCauseParams(etrace.FuncParams{
				"token": token.RefreshToken,
			})
		return nil, eerror.Err(err).
			Stack(tr).
			Err()
	}
	return &authv1.Empty{}, nil
}
