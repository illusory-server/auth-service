package grpcMapper

import (
	authv1 "github.com/illusory-server/auth-service/gen/auth"
	authUseCase "github.com/illusory-server/auth-service/internal/app/usecases/auth_usecase"
)

type AuthMapper struct{}

func (a *AuthMapper) AuthResultToAuthResponse(data *authUseCase.AuthResult) *authv1.AuthResponse {
	return &authv1.AuthResponse{
		User: &authv1.ResponseUser{
			Id:        string(data.User.Id),
			Login:     data.User.Login,
			Email:     data.User.Email,
			Role:      data.User.Role,
			IsBanned:  data.User.IsBanned,
			BanReason: *data.User.BanReason,
		},
		Tokens: &authv1.JwtTokens{
			AccessToken:  data.Tokens.AccessToken,
			RefreshToken: data.Tokens.RefreshToken,
		},
	}
}
