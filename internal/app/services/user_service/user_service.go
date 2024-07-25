package userService

import (
	"context"
	appDto "github.com/illusory-server/auth-service/internal/app/app_dto"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/internal/domain/repository"
)

type (
	Service interface {
		Create(ctx context.Context, data *appDto.RegistrationData) (*model.User, error)
	}

	service struct {
		log            domain.Logger
		userRepository repository.UserRepository
	}
)

func New(logger domain.Logger, userRepository repository.UserRepository) Service {
	return &service{
		log:            logger,
		userRepository: userRepository,
	}
}
