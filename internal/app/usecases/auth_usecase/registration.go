package authUseCase

import (
	"context"
	appDto "github.com/illusory-server/auth-service/internal/app/app_dto"
	tokenService "github.com/illusory-server/auth-service/internal/app/services/token_service"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/constants"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/eerr"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (u *useCase) Registration(ctx context.Context, data *appDto.CreateUser) (*AuthResult, error) {
	candidate, err := u.userRepository.HasByLogin(ctx, data.Login)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("login", data.Login).
			Msg("error has user candidate by login")

		return nil, eerr.Wrap(err, "[useCase.registration] userRepository.HasByLogin")
	}
	if candidate {
		u.log.Warn(ctx).
			Str("login", data.Login).
			Msg("user candidate by login")

		return nil, eerr.New(eerr.ErrConflict, "user login exist")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		u.log.Error(ctx).
			Err(err).
			Str("password", data.Password).
			Msg("failed to hash password")

		return nil, eerr.New(eerr.ErrInternal, "failed to hash password")
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
			return eerr.Wrap(err, "[useCase.Registration] userRepository.Create")
		}
		ban, err := u.banRepository.Create(txCtx, &model.Ban{
			Id:        id,
			IsBanned:  false,
			BanReason: "",
			UpdatedAt: now,
			CreatedAt: now,
		})
		if err != nil {
			return eerr.Wrap(err, "[useCase.Registration] banRepository.Create")
		}
		link := xid.New().String() + ".ru"
		_, err = u.activateRepository.Create(txCtx, id, link)
		if err != nil {
			return eerr.Wrap(err, "[useCase.Registration] activateRepository.Create")
		}

		tokensGen, err := u.tokenService.Generate(txCtx, tokenService.JwtUserData{
			Id:   pureUser.Id,
			Role: pureUser.Role,
		})
		if err != nil {
			return eerr.Wrap(err, "[useCase.Registration] tokenService.Generate")
		}

		_, err = u.tokenRepository.Save(txCtx, user.Id, tokensGen.RefreshToken)
		if err != nil {
			return eerr.Wrap(err, "[useCase.Registration] tokenRepository.Save")
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
		u.log.Error(ctx).
			Err(err).
			Interface("create_user_data", data).
			Msg("error in create user transaction")

		return nil, err
	}

	return &AuthResult{User: pureUser, Tokens: tokens}, nil
}
