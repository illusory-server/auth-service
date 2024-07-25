package grpcv1AuthService

import (
	"context"
	authv1 "github.com/illusory-server/auth-service/gen/auth"
)

func (a *AuthServiceServer) Logout(ctx context.Context, token *authv1.RefreshToken) (*authv1.Empty, error) {
	return nil, nil
}
