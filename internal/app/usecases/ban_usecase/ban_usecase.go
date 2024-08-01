package banUseCase

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/repository"
)

type (
	UseCase interface {
		BanUserById(ctx context.Context, userId domain.Id, reason string) error
		UnbanUserById(ctx context.Context, userId domain.Id) error
	}

	useCase struct {
		log           domain.Logger
		banRepository repository.BanRepository
	}
)

func (u useCase) BanUserById(ctx context.Context, userId domain.Id, reason string) error {
	//TODO implement me
	panic("implement me")
}

func (u useCase) UnbanUserById(ctx context.Context, userId domain.Id) error {
	//TODO implement me
	panic("implement me")
}

func New(logger domain.Logger, banRepository repository.BanRepository) UseCase {
	return &useCase{
		log:           logger,
		banRepository: banRepository,
	}
}
