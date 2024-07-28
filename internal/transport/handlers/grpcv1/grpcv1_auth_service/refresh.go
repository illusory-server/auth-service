package grpcv1AuthService

import (
	"context"
	authv1 "github.com/illusory-server/auth-service/gen/auth"
	errorgrpc "github.com/illusory-server/auth-service/internal/transport/errors/error_grpc"
)

func (a *AuthServiceServer) Refresh(ctx context.Context, token *authv1.RefreshToken) (*authv1.AuthResponse, error) {
	authRes, err := a.authUseCase.Refresh(ctx, token.RefreshToken)
	if err != nil {
		return nil, errorgrpc.Catch(err)
	}
	return authMapper.AuthResultToAuthResponse(authRes), nil
}
