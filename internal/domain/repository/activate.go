package repository

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
)

type ActivateRepository interface {
	Create(ctx context.Context, userId domain.Id, link string) (*model.Activate, error)

	GetByUserId(ctx context.Context, userId domain.Id) (*model.Activate, error)

	DeleteById(ctx context.Context, userId domain.Id) error

	Update(ctx context.Context, activate *model.Activate) (*model.Activate, error)
	ActivateUserById(ctx context.Context, userId domain.Id) error
	ActivateUserByLink(ctx context.Context, link string) error

	IsActivateById(ctx context.Context, userId domain.Id) (bool, error)
	HasById(ctx context.Context, userId domain.Id) (bool, error)
}
