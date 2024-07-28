package authUseCase

import (
	"context"
	appMapper "github.com/illusory-server/auth-service/internal/app/app_mapper"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/eerr"
)

func (u *useCase) Refresh(ctx context.Context, refreshToken string) (*AuthResult, error) {
	if refreshToken == "" {
		u.log.Error(ctx).
			Msg("refresh token is empty")

		return nil, eerr.New(eerr.ErrUnauthorized, eerr.MsgUnauthorized)
	}

	jwtUserData, err := u.tokenService.ValidateRefreshToken(ctx, refreshToken)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("refresh_token", refreshToken).
			Msg("invalid refresh token")
		return nil, eerr.Wrap(err, "[useCase] tokenService.ValidateRefreshToken")
	}
	has, err := u.tokenRepository.HasByValue(ctx, refreshToken)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("refresh_token", refreshToken).
			Msg("has token error")

		return nil, eerr.Wrap(err, "[useCase] tokenRepository.HasByValue")
	}
	if !has {
		u.log.Error(ctx).
			Str("refresh_token", refreshToken).
			Msg("token not found")

		return nil, eerr.New(eerr.ErrUnauthorized, eerr.MsgUnauthorized)
	}
	userDb, err := u.userRepository.GetById(ctx, jwtUserData.Id)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("id", string(jwtUserData.Id)).
			Msg("get user error")

		return nil, eerr.Wrap(err, "[useCase] userRepository.GetById")
	}

	tokens, err := u.tokenService.Generate(ctx, *jwtUserData)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Interface("data", jwtUserData).
			Msg("generate tokens error")
		return nil, eerr.Wrap(err, "[useCase] tokenService.Generate")
	}
	_, err = u.tokenRepository.Save(ctx, userDb.Id, tokens.RefreshToken)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("id", string(userDb.Id)).
			Str("refresh_token", refreshToken).
			Msg("save tokens error")
		return nil, eerr.Wrap(err, "[useCase] tokenRepository.Save")
	}

	mapper := &appMapper.UserMapper{}
	pureUser := mapper.ToPureUser(userDb, &model.Ban{})
	return &AuthResult{
		User:   pureUser,
		Tokens: tokens,
	}, nil
}
