package grpcv1AuthService

import (
	"context"
	authv1 "github.com/illusory-server/auth-service/gen/auth"
)

func (a *AuthServiceServer) Refresh(ctx context.Context, token *authv1.RefreshToken) (*authv1.AuthResponse, error) {
	//authRes, err := a.authUseCase.Refresh(ctx, token.RefreshToken)
	//if err != nil {
	//	return nil, errorgrpc.Catch(err)
	//}
	//mapper := grpcMapper.AuthMapper{}
	//return mapper.AuthResultToAuthResponseV1(authRes), nil
	return nil, nil
}
