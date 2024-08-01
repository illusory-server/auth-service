package authUseCase

import (
	"context"
	tokenService "github.com/illusory-server/auth-service/internal/app/services/token_service"
)

func (u *useCase) CheckAuth(ctx context.Context, accessToken string) (*tokenService.JwtUserData, error) {
	return u.tokenService.ValidateAccessToken(ctx, accessToken)
}
