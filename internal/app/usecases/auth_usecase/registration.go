package authUseCase

import (
	"context"
	appDto "github.com/illusory-server/auth-service/internal/app/app_dto"
	appMapper "github.com/illusory-server/auth-service/internal/app/app_mapper"
	tokenService "github.com/illusory-server/auth-service/internal/app/services/token_service"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/eerr"
)

func (u *useCase) Registration(ctx context.Context, data *appDto.RegistrationData) (*AuthResult, error) {
	user, err := u.userService.Create(ctx, data)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Interface("data", data).
			Msg("cannot create user")

		return nil, eerr.Wrap(err, "[useCase] userService.Create")
	}

	tokens, err := u.tokenService.Generate(ctx, tokenService.JwtUserData{
		Id:   user.Id,
		Role: user.Role,
	})
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Interface("data", data).
			Msg("cannot generate tokens")

		return nil, eerr.Wrap(err, "[useCase] tokenService.Generate")
	}
	_, err = u.tokenRepository.Save(ctx, user.Id, tokens.RefreshToken)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Interface("data", data).
			Msg("cannot save tokens")

		return nil, eerr.Wrap(err, "[useCase] tokenRepository.Save")
	}

	mapper := appMapper.UserMapper{}
	pureUser := mapper.ToPureUser(user, &model.Ban{
		IsBanned:  false,
		BanReason: "",
	})
	return &AuthResult{User: pureUser, Tokens: tokens}, nil
}
