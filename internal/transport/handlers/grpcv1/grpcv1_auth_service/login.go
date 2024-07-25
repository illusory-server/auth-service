package grpcv1AuthService

import (
	"context"
	authv1 "github.com/illusory-server/auth-service/gen/auth"
)

func (a *AuthServiceServer) Login(ctx context.Context, data *authv1.LoginRequest) (*authv1.AuthResponse, error) {
	//authRes, err := a.authUseCase.Login(ctx, &appDto.LoginData{
	//	Login:    data.Login,
	//	Password: data.Password,
	//})
	//if err != nil {
	//	return nil, errorgrpc.Catch(err)
	//}
	//mapper := grpcMapper.AuthMapper{}
	//return mapper.AuthResultToAuthResponseV1(authRes), nil
	return nil, nil
}
