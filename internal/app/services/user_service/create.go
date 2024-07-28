package userService

import (
	"context"
	appDto "github.com/illusory-server/auth-service/internal/app/app_dto"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/constants"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/pkg/eerr"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (s *service) Create(ctx context.Context, data *appDto.RegistrationData) (*model.User, error) {
	candidate, err := s.userRepository.HasByLogin(ctx, data.Login)
	if err != nil {
		s.log.Error(ctx).
			Err(err).
			Interface("data", data).
			Msg("failed to check user by login")

		return nil, eerr.Wrap(err, "[service.Create] userRepository.HasByLogin")
	}
	if candidate {
		s.log.Warn(ctx).
			Err(err).
			Interface("data", data).
			Msg("user login exist")

		return nil, eerr.New(eerr.ErrConflict, "user login exist")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Error(ctx).
			Err(err).
			Str("password", data.Password).
			Msg("failed to hash password")

		return nil, eerr.New(eerr.ErrInternal, "failed to hash password")
	}

	newUser, err := s.userRepository.Create(ctx, &model.User{
		Id:        domain.Id(xid.New().String()),
		Login:     data.Login,
		Email:     data.Email,
		Password:  string(hashPassword),
		Role:      constants.RoleUser,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		s.log.Error(ctx).
			Err(err).
			Interface("data", data).
			Msg("failed to create user")

		return nil, eerr.Wrap(err, "[service.Create] failed to create user")
	}
	return newUser, nil
}
