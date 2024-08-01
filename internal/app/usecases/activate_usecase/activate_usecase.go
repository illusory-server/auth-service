package activateUseCase

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/repository"
)

type (
	UseCase interface {
		LinkActivate(context.Context, string) error
		IsActivated(context.Context, domain.Id) (bool, error)
	}

	useCase struct {
		log                domain.Logger
		activateRepository repository.ActivateRepository
	}
)

func (u useCase) LinkActivate(ctx context.Context, link string) error {
	//TODO implement me
	panic("implement me")
}

func (u useCase) IsActivated(ctx context.Context, id domain.Id) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func New(logger domain.Logger, activateRepository repository.ActivateRepository) UseCase {
	return &useCase{
		log:                logger,
		activateRepository: activateRepository,
	}
}
