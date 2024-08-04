package authUseCase

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	appDto "github.com/illusory-server/auth-service/internal/app/app_dto"
	tokenService "github.com/illusory-server/auth-service/internal/app/services/token_service"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/constants"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/etrace"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (u *useCase) Registration(ctx context.Context, data *appDto.CreateUser) (*AuthResult, error) {
	candidate, err := u.userRepository.HasByLogin(ctx, data.Login)
	if err != nil {
		tr := traceUseCase.OfName("Registration").
			OfCauseFunc("userRepository.HasByLogin").
			OfCauseParams(etrace.FuncParams{
				"login": data.Login,
			})
		return nil, eerror.Err(err).
			Stack(tr).
			Err()
	}
	if candidate {
		tr := traceUseCase.OfName("Registration").
			OfCauseFunc("userRepository.HasByLogin").
			OfCauseParams(etrace.FuncParams{
				"login": data.Login,
			})
		return nil, eerror.New(eerror.MsgNotFound).
			Code(eerror.ErrNotFound).
			Stack(tr).
			Err()
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		tr := traceUseCase.OfName("Registration").
			OfCauseFunc("bcrypt.GenerateFromPassword").
			OfCauseParams(etrace.FuncParams{
				"password": data.Password,
			})
		return nil, eerror.Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}

	pureUser := &appDto.PureUser{}
	tokens := &appDto.JwtTokens{}
	now := time.Now()
	err = u.tx.WithinTransaction(ctx, func(txCtx context.Context) error {
		id := domain.Id(xid.New().String())
		user, err := u.userRepository.Create(txCtx, &model.User{
			Id:        id,
			Login:     data.Login,
			Email:     data.Email,
			Role:      constants.RoleUser,
			Password:  string(hashPassword),
			UpdatedAt: now,
			CreatedAt: now,
		})
		if err != nil {
			tr := traceUseCase.OfName("Registration").
				OfCauseFunc("userRepository.Create").
				OfCauseParams(etrace.FuncParams{
					"user": user,
				})
			return eerror.Err(err).
				Stack(tr).
				Err()
		}
		ban, err := u.banRepository.Create(txCtx, &model.Ban{
			Id:        id,
			IsBanned:  false,
			BanReason: "",
			UpdatedAt: now,
			CreatedAt: now,
		})
		if err != nil {
			tr := traceUseCase.OfName("Registration").
				OfCauseFunc("banRepository.Create").
				OfCauseParams(etrace.FuncParams{
					"ban": ban,
				})
			return eerror.Err(err).
				Stack(tr).
				Err()
		}
		// TODO - Продумать создание ссылок
		link := xid.New().String() + ".ru"
		_, err = u.activateRepository.Create(txCtx, id, link)
		if err != nil {
			tr := traceUseCase.OfName("Registration").
				OfCauseFunc("banRepository.Create").
				OfCauseParams(etrace.FuncParams{
					"ban": ban,
				})
			return eerror.Err(err).
				Stack(tr).
				Err()
		}

		jwtData := tokenService.JwtUserData{
			Id:   pureUser.Id,
			Role: pureUser.Role,
		}
		tokensGen, err := u.tokenService.Generate(txCtx, jwtData)
		if err != nil {
			tr := traceUseCase.OfName("Registration").
				OfCauseFunc("tokenService.Generate").
				OfCauseParams(etrace.FuncParams{
					"jwt_data": jwtData,
				})
			return eerror.Err(err).
				Stack(tr).
				Err()
		}

		_, err = u.tokenRepository.Save(txCtx, user.Id, tokensGen.RefreshToken)
		if err != nil {
			tr := traceUseCase.OfName("Registration").
				OfCauseFunc("tokenRepository.Save").
				OfCauseParams(etrace.FuncParams{
					"id":    user.Id,
					"token": tokensGen.RefreshToken,
				})
			return eerror.Err(err).
				Stack(tr).
				Err()
		}

		tokens.AccessToken = tokensGen.AccessToken
		tokens.RefreshToken = tokensGen.RefreshToken

		pureUser.Id = user.Id
		pureUser.Login = user.Login
		pureUser.Email = user.Email
		pureUser.Role = user.Role
		pureUser.IsBanned = ban.IsBanned
		pureUser.BanReason = &ban.BanReason
		pureUser.IsActivate = false

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &AuthResult{User: pureUser, Tokens: tokens}, nil
}
