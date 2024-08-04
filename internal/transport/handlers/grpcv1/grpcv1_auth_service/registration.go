package grpcv1AuthService

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	authv1 "github.com/illusory-server/auth-service/gen/auth"
	appDto "github.com/illusory-server/auth-service/internal/app/app_dto"
	"github.com/illusory-server/auth-service/pkg/etrace"
)

func (a *AuthServiceServer) Registration(ctx context.Context, data *authv1.RegistrationRequest) (*authv1.AuthResponse, error) {
	createData := &appDto.CreateUser{
		Login:    data.Login,
		Password: data.Password,
		Email:    data.Email,
	}
	authRes, err := a.authUseCase.Registration(ctx, createData)
	if err != nil {
		tr := traceAuthService.OfName("Refresh").
			OfCauseMethod("authUseCase.Registration").
			OfCauseParams(etrace.FuncParams{
				"data": createData,
			})
		return nil, eerror.Err(err).
			Stack(tr).
			Err()
	}

	return authMapper.AuthResultToAuthResponse(authRes), nil
}
