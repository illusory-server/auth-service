package grpcv1AuthService

import (
	"context"
	authv1 "github.com/illusory-server/auth-service/gen/auth"
)

func (a *AuthServiceServer) CheckAuth(ctx context.Context, token *authv1.AccessToken) (*authv1.JwtUser, error) {
	//TODO implement me
	panic("implement me")
}
