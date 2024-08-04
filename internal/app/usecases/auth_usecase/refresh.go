package authUseCase

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	appMapper "github.com/illusory-server/auth-service/internal/app/app_mapper"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/etrace"
)

func (u *useCase) Refresh(ctx context.Context, refreshToken string) (*AuthResult, error) {
	if refreshToken == "" {
		tr := traceUseCase.OfName("Refresh").
			OfParams(etrace.FuncParams{
				"refresh_token": refreshToken,
			})
		return nil, eerror.New(eerror.MsgUnauthorized).
			Code(eerror.ErrUnauthorized).
			Stack(tr).
			Err()
	}

	jwtUserData, err := u.tokenService.ValidateRefreshToken(ctx, refreshToken)
	if err != nil {
		tr := traceUseCase.OfName("Refresh").
			OfCauseMethod("tokenService.ValidateRefreshToken").
			OfCauseParams(etrace.FuncParams{
				"refresh_token": refreshToken,
			})
		return nil, eerror.Err(err).
			Stack(tr).
			Err()
	}
	has, err := u.tokenRepository.HasByValue(ctx, refreshToken)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("refresh_token", refreshToken).
			Msg("has token error")
		tr := traceUseCase.OfName("Refresh").
			OfCauseMethod("tokenRepository.HasByValue").
			OfCauseParams(etrace.FuncParams{
				"refresh_token": refreshToken,
			})
		return nil, eerror.Err(err).
			Stack(tr).
			Err()
	}
	if !has {
		tr := traceUseCase.OfName("Refresh").
			OfCauseMethod("tokenRepository.HasByValue").
			OfCauseParams(etrace.FuncParams{
				"refresh_token": refreshToken,
			})
		return nil, eerror.New(eerror.MsgUnauthorized).
			Code(eerror.ErrUnauthorized).
			Stack(tr).
			Err()
	}
	userDb, err := u.userRepository.GetById(ctx, jwtUserData.Id)
	if err != nil {
		tr := traceUseCase.OfName("Refresh").
			OfCauseMethod("userRepository.GetById").
			OfCauseParams(etrace.FuncParams{
				"id": jwtUserData.Id,
			})
		return nil, eerror.Err(err).
			Stack(tr).
			Err()
	}

	tokens, err := u.tokenService.Generate(ctx, *jwtUserData)
	if err != nil {
		tr := traceUseCase.OfName("Refresh").
			OfCauseMethod("tokenService.Generate").
			OfCauseParams(etrace.FuncParams{
				"jwt_data": jwtUserData,
			})
		return nil, eerror.Err(err).
			Stack(tr).
			Err()
	}
	_, err = u.tokenRepository.Save(ctx, userDb.Id, tokens.RefreshToken)
	if err != nil {
		tr := traceUseCase.OfName("Refresh").
			OfCauseMethod("tokenRepository.Save").
			OfCauseParams(etrace.FuncParams{
				"id":            userDb.Id,
				"refresh_token": refreshToken,
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
