package authUseCase

import (
	"context"
	appDto "github.com/illusory-server/auth-service/internal/app/app_dto"
	appMapper "github.com/illusory-server/auth-service/internal/app/app_mapper"
	tokenService "github.com/illusory-server/auth-service/internal/app/services/token_service"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/eerr"
	"golang.org/x/crypto/bcrypt"
)

const LoginOrPasswordIncorrect = "incorrect login or password"

func (u *useCase) Login(ctx context.Context, data *appDto.LoginData) (*AuthResult, error) {
	has, err := u.userRepository.HasByLogin(ctx, data.Login)
	if err != nil || !has {
		u.log.Error(ctx).
			Err(err).
			Interface("data", data).
			Msg("failed to check if user exists")

		return nil, eerr.New(eerr.ErrForbidden, LoginOrPasswordIncorrect)
	}

	userDb, err := u.userRepository.GetByLogin(ctx, data.Login)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Interface("data", data).
			Msg("failed get user by login")

		return nil, eerr.Wrap(err, "[useCase.Login] userRepository.GetByLogin")
	}
	isEqualPassword := bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(data.Password))
	if isEqualPassword != nil {
		u.log.Error(ctx).
			Err(err).
			Interface("data", data).
			Msg("failed to check if user exists")

		return nil, eerr.New(eerr.ErrForbidden, LoginOrPasswordIncorrect)
	}
	tokens, err := u.tokenService.Generate(ctx, tokenService.JwtUserData{Id: userDb.Id, Role: userDb.Role})
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Interface("data", data).
			Msg("cannot generate tokens")

		return nil, eerr.Wrap(err, "[useCase.Login] tokenService.Generate")
	}
	_, err = u.tokenRepository.Save(ctx, userDb.Id, tokens.RefreshToken)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Interface("data", data).
			Msg("cannot save tokens")

		return nil, eerr.Wrap(err, "[useCase.Login] tokenRepository.Save")
	}

	mapper := &appMapper.UserMapper{}
	pureUser := mapper.ToPureUser(userDb, &model.Ban{})
	return &AuthResult{
		User:   pureUser,
		Tokens: tokens,
	}, nil
}
