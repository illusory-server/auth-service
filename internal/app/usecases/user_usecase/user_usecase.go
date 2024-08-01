package userUseCase

import (
	"context"
	appDto "github.com/illusory-server/auth-service/internal/app/app_dto"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/query"
	"github.com/illusory-server/auth-service/internal/domain/repository"
)

type (
	UseCase interface {
		GetUserById(context.Context, domain.Id) (*appDto.PureUser, error)
		GetUsersByQuery(context.Context, query.PaginationQuery) ([]*appDto.PureUser, error)
		UpdateUserRole(context.Context, domain.Id, string) (*appDto.PureUser, error)
		DeleteUserById(context.Context, domain.Id) error
		CheckUserRole(context.Context, domain.Id, string) (bool, error)
	}

	useCase struct {
		log            domain.Logger
		userRepository repository.UserRepository
	}
)

func (u *useCase) GetUserById(ctx context.Context, id domain.Id) (*appDto.PureUser, error) {
	//TODO implement me
	panic("implement me")
}

func (u *useCase) GetUsersByQuery(ctx context.Context, paginationQuery query.PaginationQuery) ([]*appDto.PureUser, error) {
	//TODO implement me
	panic("implement me")
}

func (u *useCase) UpdateUserRole(ctx context.Context, id domain.Id, s string) (*appDto.PureUser, error) {
	//TODO implement me
	panic("implement me")
}

func (u *useCase) DeleteUserById(ctx context.Context, id domain.Id) error {
	//TODO implement me
	panic("implement me")
}

func (u *useCase) CheckUserRole(ctx context.Context, id domain.Id, s string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func New(log domain.Logger, userRepository repository.UserRepository) UseCase {
	return &useCase{
		log:            log,
		userRepository: userRepository,
	}
}
