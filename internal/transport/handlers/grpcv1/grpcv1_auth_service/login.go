package grpcv1AuthService

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	authv1 "github.com/illusory-server/auth-service/gen/auth"
	appDto "github.com/illusory-server/auth-service/internal/app/app_dto"
	"github.com/illusory-server/auth-service/pkg/etrace"
)

func (a *AuthServiceServer) Login(ctx context.Context, data *authv1.LoginRequest) (*authv1.AuthResponse, error) {
	dto := &appDto.LoginData{
		Login:    data.Login,
		Password: data.Password,
	}
	authRes, err := a.authUseCase.Login(ctx, dto)
	if err != nil {
		tr := traceAuthService.OfName("Login").
			OfCauseMethod("authUseCase.Login").
			OfCauseParams(etrace.FuncParams{
				"login": data.Login,
				"data":  dto,
			})
		return nil, eerror.Err(err).
			Stack(tr).
			Err()
	}
	return authMapper.AuthResultToAuthResponse(authRes), nil
}
