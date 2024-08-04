package authUseCase

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	appDto "github.com/illusory-server/auth-service/internal/app/app_dto"
	appMapper "github.com/illusory-server/auth-service/internal/app/app_mapper"
	tokenService "github.com/illusory-server/auth-service/internal/app/services/token_service"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/etrace"
	"golang.org/x/crypto/bcrypt"
)

const LoginOrPasswordIncorrect = "incorrect login or password"

func (u *useCase) Login(ctx context.Context, data *appDto.LoginData) (*AuthResult, error) {
	has, err := u.userRepository.HasByLogin(ctx, data.Login)
	if err != nil {
		tr := traceUseCase.OfName("Login").
			OfCauseMethod("userRepository.HasByLogin").
			OfCauseParams(etrace.FuncParams{
				"login": data.Login,
			})
		return nil, eerror.Err(err).
			Stack(tr).
			Err()
	}
	if !has {
		tr := traceUseCase.OfName("Login").
			OfCauseMethod("userRepository.HasByLogin").
			OfCauseParams(etrace.FuncParams{
				"login": data.Login,
			})
		return nil, eerror.Err(err).
			Code(eerror.ErrForbidden).
			Msg(LoginOrPasswordIncorrect).
			Stack(tr).
			Err()
	}

	userDb, err := u.userRepository.GetByLogin(ctx, data.Login)
	if err != nil {
		tr := traceUseCase.OfName("Login").
			OfCauseMethod("userRepository.GetByLogin").
			OfCauseParams(etrace.FuncParams{
				"login": data.Login,
			})
		return nil, eerror.Err(err).
			Stack(tr).
			Err()
	}
	isEqualPassword := bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(data.Password))
	if isEqualPassword != nil {
		u.log.Warn(ctx).
			Err(err).
			Str("password", data.Password).
			Msg("compare password error")
		tr := traceUseCase.OfName("Login").
			OfCauseMethod("bcrypt.CompareHashAndPassword").
			OfCauseParams(etrace.FuncParams{
				"password": data.Password,
			})
		return nil, eerror.Err(err).
			Code(eerror.ErrForbidden).
			Msg(LoginOrPasswordIncorrect).
			Stack(tr).
			Err()
	}
	jwtData := tokenService.JwtUserData{Id: userDb.Id, Role: userDb.Role}
	tokens, err := u.tokenService.Generate(ctx, jwtData)
	if err != nil {
		tr := traceUseCase.OfName("Login").
			OfCauseMethod("tokenService.Generate").
			OfCauseParams(etrace.FuncParams{
				"jwt_data": jwtData,
			})
		return nil, eerror.Err(err).
			Stack(tr).
			Err()
	}
	_, err = u.tokenRepository.Save(ctx, userDb.Id, tokens.RefreshToken)
	if err != nil {
		tr := traceUseCase.OfName("Login").
			OfCauseMethod("tokenRepository.Save").
			OfCauseParams(etrace.FuncParams{
				"jwt_data": jwtData,
			})
		return nil, eerror.Err(err).
			Stack(tr).
			Err()
	}

	mapper := &appMapper.UserMapper{}
	pureUser := mapper.ToPureUser(userDb, &model.Ban{})
	return &AuthResult{
		User:   pureUser,
		Tokens: tokens,
	}, nil
}
