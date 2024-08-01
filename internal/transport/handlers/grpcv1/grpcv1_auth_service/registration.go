package grpcv1AuthService

import (
	"context"
	authv1 "github.com/illusory-server/auth-service/gen/auth"
	appDto "github.com/illusory-server/auth-service/internal/app/app_dto"
	errorgrpc "github.com/illusory-server/auth-service/internal/transport/errors/error_grpc"
)

func (a *AuthServiceServer) Registration(ctx context.Context, data *authv1.RegistrationRequest) (*authv1.AuthResponse, error) {
	authRes, err := a.authUseCase.Registration(ctx, &appDto.CreateUser{
		Login:    data.Login,
		Password: data.Password,
		Email:    data.Email,
	})
	if err != nil {
		return nil, errorgrpc.Catch(err)
	}

	return authMapper.AuthResultToAuthResponse(authRes), nil
}
