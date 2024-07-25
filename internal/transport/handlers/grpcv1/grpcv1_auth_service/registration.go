package grpcv1AuthService

import (
	"context"
	authv1 "github.com/illusory-server/auth-service/gen/auth"
)

func (a *AuthServiceServer) Registration(ctx context.Context, data *authv1.RegistrationRequest) (*authv1.AuthResponse, error) {
	//authRes, err := a.authUseCase.Registration(ctx, &appDto.RegistrationData{
	//	Login:    data.Login,
	//	Password: data.Password,
	//	Email:    data.Email,
	//})
	//if err != nil {
	//	return nil, errorgrpc.Catch(err)
	//}
	//mapper := grpcMapper.AuthMapper{}
	//return mapper.AuthResultToAuthResponseV1(authRes), nil
	return nil, nil
}
